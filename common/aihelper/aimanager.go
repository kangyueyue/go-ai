package aihelper

import (
	"context"
	"sync"

	"github.com/kangyueyue/go-ai/common/logger"
)

var ctx = context.Background()

// AIHelperManager ai helper 管理器
type AIHelperManager struct {
	helpers map[string]map[string]*AIHelper // 管理用户和ai helper的映射关系
	mu      sync.RWMutex
}

// NewAIHelperManager 创建ai helper 管理器
func NewAIHelperManager() *AIHelperManager {
	return &AIHelperManager{
		helpers: make(map[string]map[string]*AIHelper),
		mu:      sync.RWMutex{},
	}
}

// GetOrCreateAIHelper 获取或创建ai helper
func (m *AIHelperManager) GetOrCreateAIHelper(userName string,
	sessionID string, modelType string,
	config map[string]interface{},
) (*AIHelper, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 获取用户的helper
	userHelpers, ok := m.helpers[userName]
	if !ok {
		// create
		userHelpers = make(map[string]*AIHelper)
		m.helpers[userName] = userHelpers
	}

	// 检查会话是否存在
	helper, ok := userHelpers[sessionID]
	if ok {
		// 存在
		return helper, nil
	}

	// 创建新的helper
	factory := GetGlobalFactory()
	helper, err := factory.CreateAIHelper(ctx, modelType, sessionID, config)
	if err != nil {
		logger.Log.Errorf("create ai helper error: %v", err)
		return nil, err
	}
	// 加入
	userHelpers[sessionID] = helper
	return helper, nil
}

// GetAIHelper 获取ai helper
func (m *AIHelperManager) GetAIHelper(userName string, sessionID string) (*AIHelper, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	userHelpers, ok := m.helpers[userName]
	if !ok {
		return nil, false
	}
	helper, ok := userHelpers[sessionID]
	if !ok {
		return nil, false
	}
	return helper, true
}

// RemoveAIHelper 移除ai helper
func (m *AIHelperManager) RemoveAIHelper(userName string, sessionID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	userHelpers, ok := m.helpers[userName]
	if !ok {
		return
	}
	// del 移除
	delete(userHelpers, sessionID)

	// 如果user 没有 session，del 映射关系
	if len(userHelpers) == 0 {
		delete(m.helpers, userName)
	}
}

// GetAllSessionID 获取指定用户的所有会话ID
func (m *AIHelperManager) GetAllSessionID(userName string) []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	userHelpers, ok := m.helpers[userName]
	if !ok {
		return nil
	}
	sessionIds := make([]string, 0, len(userHelpers))
	for sessionID := range userHelpers {
		sessionIds = append(sessionIds, sessionID)
	}
	return sessionIds
}

// 全局管理器，单例模式
var globalManager *AIHelperManager
var managerOnce sync.Once

// GetGlobalManager 获取全局管理器
func GetGlobalManager() *AIHelperManager {
	managerOnce.Do(func() {
		if globalManager == nil {
			globalManager = NewAIHelperManager()
		}
	})
	return globalManager
}
