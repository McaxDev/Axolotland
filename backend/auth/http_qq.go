package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendQQCode(c *gin.Context) {

	qq := c.Copy().Param("number")
	authcode := utils.RandomCode(6)
	expiry := time.Now().Add(10 * time.Minute)
	QQSent[qq] = MsgSentValue{
		Authcode: authcode,
		Expiry:   expiry,
	}

	c.JSON(200, utils.Resp("请求验证码成功", authcode))
}
