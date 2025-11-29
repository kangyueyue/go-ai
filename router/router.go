package router

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	enterRouter := r.Group("/api/v1")
	{
		RegisterUserRouter(enterRouter.Group("/user"))
	}
	// TODO 后续登录的接口需要jwt鉴权
	return r
}
