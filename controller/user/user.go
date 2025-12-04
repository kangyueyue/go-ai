package user

import (
	mylog "github.com/kangyueyue/go-ai/common/logger"
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
		Password string `json:"password" binding:"gt=0"`
	}

	RegisterResponse struct {
		controller.Response
		Token string `json:"token,omitempty"`
	}

	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"gt=0"`
	}

	LoginResponse struct {
		controller.Response
		Token string `json:"token,omitempty"`
	}

	CaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}

	CaptchaResponse struct {
		controller.Response
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
	mylog.Log.Infof("register success,email:%s", req.Email)
	ctx.JSON(http.StatusOK, resp)
}

// Login 登录
func Login(ctx *gin.Context) {
	req := new(LoginRequest)
	resp := new(LoginResponse)

	// 接受参数
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
		return
	}
	token, c := user.Login(req.Username, req.Password)
	if c != code.CodeSuccess {
		ctx.JSON(http.StatusOK, resp.CodeOf(c))
		return
	}
	resp.Success()
	resp.Token = token
	mylog.Log.Infof("login success,username:%s", req.Username)
	ctx.JSON(http.StatusOK, resp)

}

// Captcha 验证码
func Captcha(ctx *gin.Context) {
	req := new(CaptchaRequest)
	resp := new(CaptchaResponse)

	// 接受参数
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
		return
	}
	c := user.SendCaptcha(req.Email)
	if c != code.CodeSuccess {
		ctx.JSON(http.StatusOK, resp.CodeOf(c))
		return
	}
	resp.Success()
	mylog.Log.Infof("send to email:%s success", req.Email)
	ctx.JSON(http.StatusOK, resp)
}
