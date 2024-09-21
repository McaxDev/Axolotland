package main

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
)

var emailSent = make(map[string]emailSentValue)

type MailData struct {
	Receiver string
	AuthCode string
	Expiry   string
	Location string
}

type emailSentValue struct {
	AuthCode string
	Expiry   time.Time
}

func SendEmail(email string) error {
	authcode := utils.RandomCode(6)
	expiry := time.Now().Add(10 * time.Minute)
	emailSent[email] = emailSentValue{
		AuthCode: authcode,
		Expiry:   expiry,
	}

	var region string
	_, err := restyClient.R().SetResult(region).Get(config.GeoServer)
	if err != nil {
		return err
	}

	mailBody, err := getMailBody(&MailData{
		Receiver: email,
		AuthCode: authcode,
		Expiry:   expiry.Format("2006-01-02 15:04"),
		Location: region,
	})
	if err != nil {
		return err
	}

	if err := smtp.SendMail(
		config.SMTP.Server+":"+config.SMTP.Port,
		smtp.PlainAuth("",
			config.SMTP.Mail,
			config.SMTP.Password,
			config.SMTP.Server,
		),
		config.SMTP.Mail,
		[]string{email},
		mailBody,
	); err != nil {
		return err
	}
	return nil
}

// 生成邮件内容体的函数
func getMailBody(maildata *MailData) ([]byte, error) {

	tmplPath := "/data/email.html"
	var tmpl *template.Template
	var err error

	if _, err = os.Stat(tmplPath); err == nil {
		tmpl, err = template.ParseFiles(tmplPath)
	} else {
		tmpl, err = template.New("default").Parse(`<html><body>
				{{.Receiver}}，你的验证码是 {{.AuthCode}}。
				于{{.Expiration}}前有效，IP地址位于{{.Location}}。
		</body></html>`)
	}
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.Write([]byte(
		"From: Axolotland Gaming Club <axolotland@163.com>\r\n" +
			"To: " + maildata.Receiver + "\r\n" +
			"Subject: 验证码邮件\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n",
	))

	if err := tmpl.Execute(&buffer, maildata); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
