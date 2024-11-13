package main

import (
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {

	user, err := GetUser(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(400, utils.Resp("获取用户信息失败", err))
		return
	}

	newToken, err := GetJwt(user.ID)
	if err != nil {
		c.JSON(500, utils.Resp("更新用户凭证失败", err))
		return
	}

	c.Header("Authorization", newToken)

	c.JSON(200, utils.Resp("获取成功", gin.H{
		"userId":      user.ID,
		"username":    user.Username,
		"avatar":      user.Avatar,
		"profile":     user.Profile,
		"admin":       user.Admin,
		"money":       user.Money,
		"email":       user.Email,
		"telephone":   user.Telephone,
		"bedrockName": user.BedrockName,
		"javaName":    user.JavaName,
		"dstName":     user.DstName,
	}))
}
