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
	if ok, _ = user.IsExistUser(email); ok {
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
	if err := myemail.SendCaptcha(email, username, user.UserNameMsg); err != nil {
		return "", code.CodeServerBusy
	}

	// 6. jwt token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
}
