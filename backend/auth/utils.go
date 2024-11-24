package main

import (
	"fmt"
	"net/smtp"
	"time"
)

func AuthCode(number, authcode string, data MsgSent) (bool, error) {

	msgSent, exists := data[number]
	if !exists {
		return false, fmt.Errorf("请先申请验证码：%s\n", number)
	}

	if expiry := msgSent.Expiry; time.Now().After(expiry) {
		return false, fmt.Errorf("验证码已过期%v\n", expiry)
	}

	if authcode != msgSent.Authcode {
		return false, fmt.Errorf("验证码不正确：%s\n", authcode)
	}

	return true, nil
}

func SendEmail(email string, content []byte) error {

	if err := smtp.SendMail(
		config.SMTP.Server+":"+config.SMTP.Port,
		smtp.PlainAuth("",
			config.SMTP.Mail,
			config.SMTP.Password,
			config.SMTP.Server,
		),
		config.SMTP.Mail,
		[]string{email},
		content,
	); err != nil {
		return err
	}
	return nil
}
