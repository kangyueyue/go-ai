package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai/common/code"
	"github.com/kangyueyue/go-ai/controller"
	"github.com/kangyueyue/go-ai/service/user"
)

type (
	RegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Captcha  string `json:"captcha"`
		Password string `json:"password" binding:"gte=0"`
	}

	RegisterResponse struct {
		controller.Response
		Token string `json:"token,omitempty"`
	}
)

// Register 注册
func Register(ctx *gin.Context) {
	req := new(RegisterRequest)
	resp := new(RegisterResponse)

	// 接受参数
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
		return
	}
	token, c := user.Register(req.Email, req.Password, req.Captcha)
	if c != code.CodeSuccess {
		ctx.JSON(http.StatusOK, resp.CodeOf(c))
		return
	}
	resp.Success()
	resp.Token = token
	ctx.JSON(http.StatusOK, resp)
}

// Login 登录
func Login(ctx *gin.Context) {
}

// Captcha 验证码
func Captcha(ctx *gin.Context) {

}
