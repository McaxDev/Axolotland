package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthJwtMid(
	logicFunc func(user *User, c *gin.Context),
) func(c *gin.Context) {
	return func(c *gin.Context) {

		user, err := GetUser(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(400, utils.Resp("用户验证失败", err, nil))
			return
		}
		logicFunc(user, c)
	}
}

func AuthCaptchaMid(
	logicFunc func(c *gin.Context),
) func(c *gin.Context) {
	return func(c *gin.Context) {

		var request struct {
			CaptchaID    string
			CaptchaValue string
		}
		if err := utils.LoadBodyFromCtx(c, &request); err != nil {
			c.JSON(400, utils.Resp("用户请求有误", err, nil))
			return
		}

		CaptchaResponse, err := AuthClient.Auth(
			context.Background(),
			&rpc.Authcode{
				Codetype: "captcha",
				Number:   request.CaptchaID,
				Authcode: request.CaptchaValue,
			},
		)
		if err != nil || !CaptchaResponse.Data {
			c.JSON(400, utils.Resp("人机验证失败", err, nil))
			return
		}

		logicFunc(c)
	}
}
