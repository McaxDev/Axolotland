package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
)

func Signout(user *User, c *gin.Context) {

	var request rpc.Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	response, err := VerifyClient.VerifyCode(
		context.Background(), &request,
	)
	if err != nil || !response.Success {
		c.JSON(400, utils.Resp("联系方式验证失败", err))
		return
	}

	if err := DB.Delete(&user).Error; err != nil {
		c.JSON(500, utils.Resp("注销删除失败", err))
		return
	}

	c.JSON(200, utils.Resp("注销成功，感谢使用", nil))
}
