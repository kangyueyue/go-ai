package message

import (
	"github.com/kangyueyue/go-ai/common/mysql"
	"github.com/kangyueyue/go-ai/model"
)

// 获取所有消息
func GetAllMessages() ([]model.Message, error) {
	var msgs []model.Message
	err := mysql.DB.Order("created_at asc").Find(&msgs).Error
	return msgs, err
}

// CreateMessage 创建消息
func CreateMessage(message *model.Message) (*model.Message, error) {
	err := mysql.DB.Create(message).Error
	return message, err
}
