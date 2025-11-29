package myemail

import (
	"fmt"

	"github.com/kangyueyue/go-ai/config"
	"gopkg.in/gomail.v2"
)

// SendCaptcha 发送验证码
func SendCaptcha(email, code, msg string) error {
	m := gomail.NewMessage()

	// 发件人
	m.SetHeader("From", config.GetConfig().EmailConfig.Email)
	// 收件人
	m.SetHeader("To", email)
	// 主题
	m.SetHeader("Subject", "来自GopherAI的信息")
	// 正文内容（纯文本形式，也可以用 text/html）
	m.SetBody("text/plain", msg+" "+code)

	// 配置 SMTP 服务器和授权码,587：是 SMTP 的明文/STARTTLS 端口号
	d := gomail.NewDialer("smtp.qq.com", 587, config.GetConfig().EmailConfig.Email, config.GetConfig().EmailConfig.AuthCode)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:\n", err)
		return err
	}
	fmt.Printf("send mail success\n")
	return nil
}
