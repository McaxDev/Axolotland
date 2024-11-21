package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/McaxDev/Axolotland/backend/gameapi/config"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthJwtMiddle(
	logicFunc func(user *rpc.JwtResponse, c *gin.Context),
) func(c *gin.Context) {
	return func(c *gin.Context) {

		user, err := config.AccountClient.VerifyJwt(
			context.Background(),
			&rpc.JwtRequest{JWT: c.GetHeader("Authorization")},
		)
		if err != nil {
			c.JSON(500, utils.Resp("身份验证服务器错误", err))
			return
		}

		logicFunc(user, c)
	}
}
