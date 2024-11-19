package main

import (
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(user *User, c *gin.Context) {

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

func SetUserInfo(user *User, c *gin.Context) {

	var request struct {
		Type  string
		Value string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err))
		return
	}

	query := DB.Model(&user)

	switch request.Type {
	case "avatar":
		query = query.Update("Avatar", request.Value)
	case "profile":
		query = query.Update("Profile", request.Value)
	default:
		c.JSON(400, utils.Resp("不支持设置此种信息", nil))
		return
	}

	if err := query.Error; err != nil {
		c.JSON(500, utils.Resp("资料更新失败", err))
		return
	}

	c.JSON(200, utils.Resp("资料更新成功", nil))
}
