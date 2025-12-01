package user

import (
	"errors"

	"github.com/kangyueyue/go-ai/common/mysql"
	"github.com/kangyueyue/go-ai/model"
	"github.com/kangyueyue/go-ai/utils"
	"gorm.io/gorm"
)

// IsExistUserByEmail 判断用户是否存在
func IsExistUserByEmail(email string) (bool, *model.User) {
	user, err := mysql.GetUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
		// 不存在
		return false, nil
	}
	return true, user // 存在
}

// IsExistUserByUsername 判断用户是否存在
func IsExistUserByUsername(username string) (bool, *model.User) {
	user, err := mysql.GetUserByUserName(username)
	if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
		// 不存在
		return false, nil
	}
	return true, user // 存在
}

// Register  注册
func Register(email, password, username string) (*model.User, bool) {
	if user, err := mysql.InsertUser(&model.User{
		Email:    email,
		Name:     username,
		Password: utils.MD5(password), // 加密存储
		Username: username,
	}); err != nil {
		return nil, false
	} else {
		return user, true
	}
}
