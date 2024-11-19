package handler

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/GameAPI/config"
	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func GameBackup(c *gin.Context) {

	user, err := config.AccountClient.VerifyJwt(
		context.Background(),
		&rpc.JwtRequest{JWT: c.GetHeader("Authorization")},
	)
	if err != nil || !user.Admin {
		c.JSON(400, utils.Resp("身份验证失败", err))
		return
	}

	var request struct {
		Server string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("读取请求参数失败", err))
		return
	}

	var server *config.Server
	for index := range config.Servers {
		if temp := &config.Servers[index]; temp.Name == request.Server {
			server = temp
		}
	}

}
