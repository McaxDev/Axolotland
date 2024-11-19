package main

import (
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthJwtMiddle(
	logicFunc func(user *User, c *gin.Context),
) func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := GetUser(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(400, utils.Resp("用户验证失败", err))
			return
		}
		logicFunc(user, c)
	}
}
