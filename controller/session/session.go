package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai/common/code"
	"github.com/kangyueyue/go-ai/common/logger"
	"github.com/kangyueyue/go-ai/controller"
	"github.com/kangyueyue/go-ai/model"
	"github.com/kangyueyue/go-ai/service/session"
)

type (
	GetUserSessionsByUserNameReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
	}
	GetUserSessionsByUserNameResp struct {
		controller.Response
		Sessions []model.SessionInfo `json:"sessions",omitempty`
	}
	ChatSendReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
		SessionID    string `json:"sessionId,omitempty" binding:"required"`
	}
	ChatSendResp struct {
		controller.Response
		// AI回答
		AiInformation string `json:"Information,omitempty"`
	}
	ChatHistoryReq struct {
		SessionID string `json:"sessionId,omitempty" binding:"required"`
	}
	ChatHistoryResp struct {
		controller.Response
		History []model.History `json:"history"`
	}
	CreateSessionAndSendMessageReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
	}
	CreateSessionAndSendMessageResp struct {
		controller.Response
		// AI回答
		AiInformation string `json:"Information,omitempty"`
		// 当前会话ID
		SessionID string `json:"sessionId,omitempty"`
	}
)

// GetUserSessionsByUserName 获取用户的会话列表
func GetUserSessionsByUserName(c *gin.Context) {
	res := new(GetUserSessionsByUserNameResp)
	userName := c.GetString("userName") // from JWT middleware
	userSessions, err := session.GetUserSessionsByUserName(userName)

	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	res.Sessions = userSessions
	c.JSON(http.StatusOK, res)
}

// ChatSend 聊天发送消息
func ChatSend(c *gin.Context) {
	res := new(ChatSendResp)
	req := new(ChatSendReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	logger.Log.Infof("userName:%s", userName)
	aiInformation, code_ := session.ChatSend(userName, req.UserQuestion,
		req.ModelType, req.SessionID)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success() // success
	res.AiInformation = aiInformation
	c.JSON(http.StatusOK, res)
}

// CreateSessionAndSendMessage 创建会话并发送消息
func CreateSessionAndSendMessage(c *gin.Context) {
	req := new(CreateSessionAndSendMessageReq)
	res := new(CreateSessionAndSendMessageResp)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	logger.Log.Infof("userName:%s", userName)
	session_id, aiInformation, code_ := session.CreateSessionAndSendMessage(userName, req.UserQuestion, req.ModelType)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success()
	res.AiInformation = aiInformation
	res.SessionID = session_id
	c.JSON(http.StatusOK, res)
}

// CreateStreamSessionAndSendMessage 创建流式会话并发送消息
func CreateStreamSessionAndSendMessage(c *gin.Context) {

}

// ChatHistory 聊天历史记录
func ChatHistory(c *gin.Context) {
	req := new(ChatHistoryReq)
	res := new(ChatHistoryResp)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	history, code_ := session.ChatHistory(userName, req.SessionID)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success()
	res.History = history
	c.JSON(http.StatusOK, res)
}

// ChatStreamSend 流式聊天发送消息
func ChatStreamSend(c *gin.Context) {

}
