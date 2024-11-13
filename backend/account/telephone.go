package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
)

func SetTelephone(c *gin.Context) {

	var request struct {
		Telephone string
		Authcode  string
	}

	user, err := GetUser(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(400, utils.Resp("读取用户信息失败", err))
		return
	}

	SMSResponse, err := VerifyClient.VerifyCode(
		context.Background(), &rpc.Request{
			Codetype: "sms",
			Number:   request.Telephone,
			Authcode: request.Authcode,
		},
	)

	if err != nil || !SMSResponse.Success {
		c.JSON(400, utils.Resp("手机号码验证失败", err))
		return
	}

	if err := DB.Model(&user).Update(
		"Telephone", request.Telephone,
	).Error; err != nil {
		c.JSON(500, utils.Resp("手机号码修改失败", err))
		return
	}

	c.JSON(200, utils.Resp("手机号码修改成功", nil))
}
