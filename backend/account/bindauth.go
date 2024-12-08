package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindAuth(user *User, c *gin.Context) {

	var request rpc.Authcode

	response, err := AuthClient.Auth(
		context.Background(), &request,
	)

	if err != nil || !response.Data {
		c.JSON(400, utils.Resp("号码验证失败", err))
		return
	}

	query := DB.Model(&user)

	if request.Codetype == "telephone" {
		query = query.Update("Telephone", request.Number)
	} else {
		query = query.Update("Email", request.Number)
	}

	if err := query.Error; err != nil {
		c.JSON(500, utils.Resp("号码修改失败", err))
		return
	}

	c.JSON(200, utils.Resp("号码修改成功", nil))
}
