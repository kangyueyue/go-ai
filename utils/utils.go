package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/kangyueyue/go-ai/model"
)

// GetRandomNumbers 生成随机name
func GetRandomNumbers(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < num; i++ {
		// 0-9 随机数
		digit := r.Intn(10)
		code += strconv.Itoa(digit)
	}
	return code
}

// MD5 加密
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// ConvertToSchemaMessages 转换为schema消息
func ConvertToSchemaMessages(msgs []*model.Message) []*schema.Message {
	schemaMsg := make([]*schema.Message, 0, len(msgs))
	for _, m := range msgs {
		role := schema.Assistant
		if !m.IsUser {
			role = schema.User
		}
		schemaMsg = append(schemaMsg, &schema.Message{
			Role:    role,
			Content: m.Content,
		})
	}
	return schemaMsg
}

// ConvertToModelMessages 转换为model消息
func ConvertToModelMessages(sessionId string, userName string, msg *schema.Message) *model.Message {
	return &model.Message{
		SessionID: sessionId,
		UserName:  userName,
		Content:   msg.Content,
	}
}
