package main

import (
	"fmt"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func HttpHandler(c *gin.Context) {

	codetype := c.Param("codetype")

	switch codetype {
	case "email":
		email := c.Param("number")
		if err := SendEmail(email); err != nil {
			c.JSON(500, utils.Resp("邮件发送失败", nil))
			fmt.Println(err.Error())
		}
		c.JSON(200, utils.Resp("验证码发送成功", nil))
	case "sms":
		telephone := c.Param("number")
		if err := SendSMS(telephone); err != nil {
			c.JSON(500, utils.Resp("短信发送失败", nil))
			fmt.Println(err.Error())
		}
		c.JSON(200, utils.Resp("验证码发送成功", nil))
	case "captcha":
		SendCaptcha(c)
	default:
		c.JSON(400, utils.Resp("路径不存在", nil))
		return
	}
}
