package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResetPassword(c *gin.Context) {

	var request struct {
		Codetype string
		Number   string
		Authcode string
		Password string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	response, err := VerifyClient.VerifyCode(
		context.Background(),
		&rpc.Request{
			Codetype: request.Codetype,
			Number:   request.Number,
			Authcode: request.Authcode,
		},
	)
	if err != nil || !response.Success {
		c.JSON(400, utils.Resp("验证失败", err))
		return
	}

	numberType := "telephone"
	if request.Codetype == "email" {
		numberType = "email"
	}

	if err := DB.Model(new(User)).Where(
		numberType+" = ?", request.Number,
	).Update(
		"Password", request.Password,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("不存在这个用户", nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("密码修改失败", err))
		return
	}

	c.JSON(200, utils.Resp("密码修改成功", nil))
}
