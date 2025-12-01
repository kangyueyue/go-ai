package user

import (
	"github.com/kangyueyue/go-ai/common/code"
	myemail "github.com/kangyueyue/go-ai/common/email"
	"github.com/kangyueyue/go-ai/common/redis"
	"github.com/kangyueyue/go-ai/dao/user"
	"github.com/kangyueyue/go-ai/model"
	"github.com/kangyueyue/go-ai/utils"
	myjwt "github.com/kangyueyue/go-ai/utils/jwt"
)

// Register 注册,返回 token
func Register(email, password, captcha string) (string, code.Code) {
	var ok bool
	var UserInformation *model.User

	// 1.判断用户是否存在
	if ok, _ = user.IsExistUserByEmail(email); ok {
		return "", code.CodeUserExist
	}

	// 2. 从redis中判断验证码是否正确
	if ok, _ = redis.CheckCaptcha(email, captcha); !ok {
		return "", code.CodeInvalidCaptcha
	}

	// 3. gen username
	username := utils.GetRandomNumbers(11)

	// 4. store
	if UserInformation, ok = user.Register(email, password, username); !ok {
		return "", code.CodeServerBusy
	}

	// 5.账户发送到邮箱账号，之后凭借账户登入
	if err := myemail.SendCaptcha(email, username, myemail.UserNameMsg); err != nil {
		return "", code.CodeServerBusy
	}

	// 6. jwt token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
}

// Login login
func Login(username, password string) (string, code.Code) {
	var ok bool
	var UserInformation *model.User
	// 1.判断用户是否存在
	if ok, UserInformation = user.IsExistUserByUsername(username); !ok {
		return "", code.CodeUserNotExist
	}
	// 2.check password
	if UserInformation.Password != utils.MD5(password) {
		return "", code.CodeInvalidPassword
	}
	// 3.return token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}
	return token, code.CodeSuccess
}

// SendCaptcha 发送验证码
func SendCaptcha(email string) code.Code {
	// 创建6为随机数
	send_code := utils.GetRandomNumbers(6)
	// 1. 存放到redis中
	if err := redis.SetCaptchaForEmail(email, send_code); err != nil {
		return code.CodeServerBusy
	}
	// 2.发送邮件
	if err := myemail.SendCaptcha(email, send_code, myemail.CodeMsg); err != nil {
		return code.CodeServerBusy
	}
	return code.CodeSuccess
}
