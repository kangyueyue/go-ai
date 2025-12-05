package session

import (
	"github.com/kangyueyue/go-ai/common/mysql"
	"github.com/kangyueyue/go-ai/model"
)

// CreateSession 创建会话
func CreateSession(session *model.Session) (*model.Session, error) {
	err := mysql.DB.Create(session).Error
	return session, err
}
