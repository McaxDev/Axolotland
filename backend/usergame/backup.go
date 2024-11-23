package main

import (
	"context"

	accountrpc "github.com/McaxDev/Axolotland/backend/account/rpc"
	gameapirpc "github.com/McaxDev/Axolotland/backend/gameapi/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func WorldBackup(user *accountrpc.JwtResponse, c *gin.Context) {

	var request struct {
		Server string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("读取请求参数失败", err))
		return
	}

	response, err := GameapiClient.WorldBackup(
		context.Background(),
		&gameapirpc.BackupRequest{Server: request.Server},
	)
	if err != nil || !response.Success {
		c.JSON(500, utils.Resp("备份失败", err))
		return
	}

	c.JSON(200, utils.Resp("备份成功", nil))
}
