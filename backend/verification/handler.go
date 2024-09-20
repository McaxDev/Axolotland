package main

import "github.com/gin-gonic/gin"

func HttpHandler(c *gin.Context) {

	codetype := c.Param("codetype")

	switch codetype {
	case "email":
		email := c.Param("number")
		SendEmail(email)
	case "sms":
		telephone := c.Param("number")
		SendSMS(telephone)
	case "captcha":
		SendCaptcha(c)
	default:
		c.JSON(400, Response("路径不存在", nil))
		return
	}
}
