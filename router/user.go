package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai/controller/user"
)

// RegisterUserRouter 注册用户路由
func RegisterUserRouter(r *gin.RouterGroup) {
	{
		r.POST("/register", user.Register) // 注册
		r.POST("/login", user.Login)       // 登入
		r.POST("/captcha", user.Captcha)   // 验证码
	}
}
