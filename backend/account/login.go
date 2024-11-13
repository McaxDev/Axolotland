package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var request struct {
		Username     string
		Password     string
		CaptchaID    string
		CaptchaValue string
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("请求参数有误", err))
		return
	}

	captchaResponse, err := VerifyClient.VerifyCode(
		context.Background(), &rpc.Request{
			Codetype: "captcha",
			Number:   request.CaptchaID,
			Authcode: request.CaptchaValue,
		},
	)

	if err != nil || !captchaResponse.Success {
		c.JSON(400, utils.Resp("人机验证失败", err))
		return
	}

	var user User

	if err := DB.First(
		&user, "username = ?", request.Username,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("你尚未注册", nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("用户查询失败", err))
		return
	}

	if request.Password != user.Password {
		c.JSON(400, utils.Resp("密码不正确", nil))
		return
	}

	token, err := GetJwt(user.ID)
	if err != nil {
		c.JSON(500, utils.Resp("用户凭证生成失败", err))
		return
	}

	c.JSON(200, utils.Resp("登录成功", gin.H{
		"token": token,
	}))
}
