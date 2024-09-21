package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HttpHandler(c *gin.Context) {

	codetype := c.Param("codetype")

	switch codetype {
	case "email":
		email := c.Param("number")
		if err := SendEmail(email); err != nil {
			c.JSON(500, Response("邮件发送失败", nil))
			fmt.Println(err.Error())
		}
	case "sms":
		telephone := c.Param("number")
		if err := SendSMS(telephone); err != nil {
			c.JSON(500, Response("短信发送失败", nil))
			fmt.Println(err.Error())
		}
	case "captcha":
		SendCaptcha(c)
	default:
		c.JSON(400, Response("路径不存在", nil))
		return
	}
}
