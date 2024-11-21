package main

import (
	"context"
	"strings"

	"github.com/McaxDev/Axolotland/backend/GameAPI/servers"
	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func SendCmd(user *rpc.JwtResponse, c *gin.Context) {

	var request struct {
		Server  string
		Command string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	server := servers.GetServerByName(request.Server)
	if server == nil {
		c.JSON(400, utils.Resp("不存在这个服务器", nil))
		return
	}

	cmdName := strings.Split(
		strings.TrimPrefix(request.Command, "/"), " ",
	)[0]
	if !user.Admin && !slices.Contains(
		server.AllowedCommands, cmdName,
	) {
		c.JSON(400, utils.Resp("你没有权限执行这个命令", nil))
		return
	}

	response, err := server.SendCmd(
		context.Background(), request.Command,
	)
	if err != nil {
		c.JSON(500, utils.Resp("命令发送失败", err))
		return
	}

	c.JSON(200, utils.Resp("命令执行成功", response))
}
