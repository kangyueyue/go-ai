package aihelper

import (
	"github.com/kangyueyue/go-ai/common/mq"
	"github.com/kangyueyue/go-ai/model"
	"sync"
)

// AIHelper ai helper
type AIHelper struct {
	model    AIModel
	messages []*model.Message
	mu       sync.RWMutex
	// 一个会话绑定一个
	sessionID string
	saveFunc  func(*model.Message) (*model.Message, error)
}

// NewAIHelper 创建一个ai helper
func NewAIHelper(model_ AIModel, sessionID string) *AIHelper {
	return &AIHelper{
		model:     model_,
		messages:  make([]*model.Message, 0),
		mu:        sync.RWMutex{},
		sessionID: sessionID,
		saveFunc: func(msg *model.Message) (*model.Message, error) {
			// 保存策略是通过m异步q保存到db中
			data := mq.GenerateMessageMQParam(msg.SessionID, msg.Content, msg.UserName, msg.IsUser)
			err := mq.RMQMessage.Publish(string(data)) // publish 生产者
			return msg, err
		},
	}
}

// SetSaveFunc 设置保存函数
func (h *AIHelper) SetSaveFunc(saveFunc func(*model.Message) (*model.Message, error)) {
	h.saveFunc = saveFunc
}
