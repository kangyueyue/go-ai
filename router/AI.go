package router

import (
	"github.com/gin-gonic/gin"
)

// AIRouter AI router
func AIRouter(r *gin.RouterGroup) {
	{
		r.GET("/chat/sessions")
		r.POST("/chat/send-new-session")
		r.POST("/chat/send")
		r.POST("/chat/history")
		// r.POST("/chat/tts", AI.ChatSpeech)                  // ChatSpeechHandler
		r.POST("/chat/send-stream-new-session")
		r.POST("/chat/send-stream")
	}
}
