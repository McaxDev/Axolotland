package main

import (
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetUsername(user *User, c *gin.Context) {

	var request struct {
		Username string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	if err := DB.First(
		new(User), "username = ?", request.Username,
	).Error; err == nil {
		c.JSON(400, utils.Resp("已经有人使用过这个名称了", nil))
		return
	}

	if err := DB.Model(&user).Update(
		"Username", request.Username,
	).Error; err != nil {
		c.JSON(500, utils.Resp("用户名更新失败", err))
		return
	}

	c.JSON(200, utils.Resp("用户名更新成功", nil))
}
