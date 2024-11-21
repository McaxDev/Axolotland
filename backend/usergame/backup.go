package main

import (
	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/McaxDev/Axolotland/backend/gameapi/servers"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func GameBackup(user *rpc.JwtResponse, c *gin.Context) {

	if !user.Admin {
		c.JSON(400, utils.Resp("你不是管理员", nil))
		return
	}

	var request struct {
		Server string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("读取请求参数失败", err))
		return
	}

	server := servers.GetServerByName(request.Server)
	if server == nil {
		c.JSON(400, utils.Resp("不存在这个服务器", nil))
		return
	}

}
