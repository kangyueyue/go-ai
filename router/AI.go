package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai/controller/session"
)

// AIRouter AI router
func AIRouter(r *gin.RouterGroup) {
	{
		r.GET("/chat/sessions", session.GetUserSessionsByUserName)
		r.POST("/chat/send-new-session", session.CreateSessionAndSendMessage)
		r.POST("/chat/send", session.ChatSend)
		r.POST("/chat/history", session.ChatHistory)
		// r.POST("/chat/tts", AI.ChatSpeech)                  // ChatSpeechHandler
		r.POST("/chat/send-stream-new-session", session.CreateStreamSessionAndSendMessage)
		r.POST("/chat/send-stream", session.ChatStreamSend)
	}
}
