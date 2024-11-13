package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var request struct {
		Username     string
		Password     string
		Email        string
		EmailCode    string
		CaptchaID    string
		CaptchaValue string
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("请求参数错误", err))
		return
	}

	if err := DB.First(
		new(User), "username = ?", request.Username,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此用户已经注册过了", nil))
		return
	}

	if err := DB.First(
		new(User), "email = ?", request.Email,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此邮箱已经注册过了", nil))
		return
	}

	MailResponse, err := VerifyClient.VerifyCode(
		context.Background(),
		&rpc.Request{
			Codetype: "email",
			Number:   request.Email,
			Authcode: request.EmailCode,
		},
	)

	if err != nil || !MailResponse.Success {
		c.JSON(400, utils.Resp("邮箱验证失败", err))
		return
	}

	CaptchaResponse, err := VerifyClient.VerifyCode(
		context.Background(),
		&rpc.Request{
			Codetype: "captcha",
			Number:   request.CaptchaID,
			Authcode: request.CaptchaValue,
		},
	)

	if err != nil || !CaptchaResponse.Success {
		c.JSON(400, utils.Resp("人机验证失败", err))
		return
	}

	user := User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Admin:    false,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, utils.Resp("无法在数据库里创建新用户", err))
		return
	}

	token, err := GetJwt(user.ID)
	if err != nil {
		c.JSON(500, utils.Resp("JWT生成失败", err))
		return
	}

	c.JSON(200, utils.Resp("用户创建成功", gin.H{
		"token": token,
	}))
}
