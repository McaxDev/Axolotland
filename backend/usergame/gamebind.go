package main

import (
	"context"

	account "github.com/McaxDev/Axolotland/backend/account/rpc"
	gameapi "github.com/McaxDev/Axolotland/backend/gameapi/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func Gamebind(user *account.JwtResponse, c *gin.Context) {

	var request gameapi.BindRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	response, err := GameapiClient.GameBind(
		context.Background(), &request,
	)
	if err != nil || !response.Success {
		c.JSON(400, utils.Resp("绑定失败", err))
		return
	}

	c.JSON(200, utils.Resp("绑定成功", nil))
}
