package main

import (
	"context"

	accountrpc "github.com/McaxDev/Axolotland/backend/account/rpc"
	gameapirpc "github.com/McaxDev/Axolotland/backend/gameapi/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendCmd(user *accountrpc.JwtResponse, c *gin.Context) {

	var request struct {
		Server  string
		Command string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	response, err := GameapiClient.SendCmd(
		context.Background(), &gameapirpc.CmdRequest{
			Server:  request.Server,
			Command: request.Command,
			Admin:   user.Admin,
		},
	)
	if err != nil {
		c.JSON(500, utils.Resp("命令发送失败", err))
		return
	}

	c.JSON(200, utils.Resp("命令执行成功", response.Response))
}
