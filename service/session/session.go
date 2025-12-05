package session

import (
	"context"

	"github.com/google/uuid"
	"github.com/kangyueyue/go-ai/common/aihelper"
	"github.com/kangyueyue/go-ai/common/code"
	"github.com/kangyueyue/go-ai/common/logger"
	"github.com/kangyueyue/go-ai/dao/session"
	"github.com/kangyueyue/go-ai/model"
)

var ctx = context.Background()

// GetUserSessionsByUserName 获取用户的会话列表
func GetUserSessionsByUserName(userName string) ([]model.SessionInfo, error) {
	manager := aihelper.GetGlobalManager() // 获得全局的aihelper管理器
	Sessions := manager.GetAllSessionID(userName)

	var SessionInfos []model.SessionInfo

	for _, session := range Sessions {
		SessionInfos = append(SessionInfos, model.SessionInfo{
			SessionID: session,
			Title:     session, // 暂时用sessionID作为标题，后续重构需要的时候可以更改
		})
	}

	return SessionInfos, nil
}

// ChatSend 聊天发送消息
func ChatSend(userName string, userQuestion string,
	modelType string, sessionID string,
) (string, code.Code) {
	// 获取helper
	manager := aihelper.GetGlobalManager()
	config := map[string]interface{}{
		"apiKey": "your-api-key", // TODO: 从配置中获取
	}
	helper, err := manager.GetOrCreateAIHelper(userName, sessionID, modelType, config)
	if err != nil {
		return "", code.AIModelFail
	}
	// 生成ai回答
	res, err := helper.GenerateResponse(userName, ctx, userQuestion)
	if err != nil {
		logger.Log.Errorf("ChatSend GenerateResponse error: %v")
		return "", code.AIModelFail
	}
	return res.Content, code.CodeSuccess
}

// ChatHistory 聊天历史记录
func ChatHistory(userName, sessionId string) ([]model.History, code.Code) {
	// helper中的消息历史
	manager := aihelper.GetGlobalManager()
	helper, ok := manager.GetAIHelper(userName, sessionId)
	if !ok {
		return nil, code.CodeServerBusy
	}
	message := helper.GetHistory()
	history := make([]model.History, 0, len(message))

	// 转化为为历史消息
	for i, msg := range message {
		isUser := i%2 == 0
		history = append(history, model.History{
			IsUser:  isUser,
			Content: msg.Content,
		})
	}
	return history, code.CodeSuccess
}

// CreateSessionAndSendMessage 创建会话并发送消息
func CreateSessionAndSendMessage(
	userName, userQuestion, modelType string,
) (string, string, code.Code) {
	// 创建一个新的会话
	newSession := &model.Session{
		ID:       uuid.New().String(),
		UserName: userName,
		Title:    userQuestion,
	}
	createSession, err := session.CreateSession(newSession)
	if err != nil {
		logger.Log.Errorf("CreateSessionAndSendMessage CreateSession error: %v")
		return "", "", code.CodeServerBusy
	}

	manager := aihelper.GetGlobalManager()
	config := map[string]interface{}{
		"apiKey": "your-api-key", // TODO: 从配置中获取
	}
	helper, err := manager.GetOrCreateAIHelper(userName, createSession.ID, modelType, config)
	if err != nil {
		return "", "", code.AIModelFail
	}
	// 生成Ai回答
	res, err := helper.GenerateResponse(userName, ctx, userQuestion)
	if err != nil {
		logger.Log.Errorf("ChatSend GenerateResponse error: %v")
		return "", "", code.AIModelFail
	}
	return createSession.ID, res.Content, code.CodeSuccess
}
