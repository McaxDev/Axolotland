package utils

import (
	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/gin-gonic/gin"
)

func AuthJwtMiddle(
	logicFunc func(user *rpc.JwtResponse, c *gin.Context),
) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
