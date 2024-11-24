package main

import (
	"bytes"
	"html/template"
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

type MailData struct {
	Email    string
	Authcode string
	Expiry   string
	Location string
}

func SendEmailCode(c *gin.Context) {

	email := c.Param("number")
	authcode := utils.RandomCode(6)
	expiry := time.Now().Add(10 * time.Minute)
	EmailSent[email] = MsgSentValue{
		Authcode: authcode,
		Expiry:   expiry,
	}

	var region string
	_, err := RestyClient.R().SetResult(region).Get(config.GeoServer)
	if err != nil {
		c.JSON(500, utils.Resp("获取你的地理位置失败", nil))
		return
	}

	tmpl, err := template.ParseFiles("/data/email.html")
	if err != nil {
		c.JSON(500, utils.Resp("读取邮件模板失败", err))
		return
	}

	var buffer bytes.Buffer
	buffer.Write([]byte(
		"From: Axolotland Gaming Club <axolotland@163.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: 验证码邮件\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n",
	))

	if err := tmpl.Execute(&buffer, &MailData{
		Email:    email,
		Authcode: authcode,
		Expiry:   expiry.Format("2006-01-02 15:04"),
		Location: region,
	}); err != nil {
		c.JSON(500, utils.Resp("生成邮件失败", err))
		return
	}

	if err := SendEmail(email, buffer.Bytes()); err != nil {
		c.JSON(500, utils.Resp("邮件发送失败", err))
		return
	}

	c.JSON(200, utils.Resp("验证码发送成功", nil))
}
