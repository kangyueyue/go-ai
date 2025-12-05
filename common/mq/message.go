package mq

import (
	"encoding/json"

	"github.com/kangyueyue/go-ai/common/logger"
	"github.com/kangyueyue/go-ai/dao/message"
	"github.com/kangyueyue/go-ai/model"
	"github.com/streadway/amqp"
)

// MessageMQParam 消息队列参数
type MessageMQParam struct {
	SessionId string `json:"session_id"`
	Content   string `json:"content"`
	UserName  string `json:"user_name"`
	IsUser    bool   `json:"is_user"`
}

// GenerateMessageMQParam 生成消息队列参数
func GenerateMessageMQParam(sessionId, content, userName string, isUser bool) []byte {
	param := &MessageMQParam{
		SessionId: sessionId,
		Content:   content,
		UserName:  userName,
		IsUser:    isUser,
	}
	data, _ := json.Marshal(param)
	return data
}

// MqMessage 消息队列消息处理
func MqMessage(msg *amqp.Delivery) error {
	var param MessageMQParam
	err := json.Unmarshal(msg.Body, &param)
	if err != nil {
		return err
	}
	newMsg := &model.Message{
		SessionID: param.SessionId,
		Content:   param.Content,
		UserName:  param.UserName,
		IsUser:    param.IsUser,
	}
	_, err = message.CreateMessage(newMsg)
	if err == nil {
		logger.Log.Infof("use mq to save message success:%v", newMsg.Content)
	}
	return err
}
