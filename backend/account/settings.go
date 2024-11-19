package main

import (
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSettings(user *User, c *gin.Context) {

	type SettingsStruct struct {
		Name    string
		Comment string
		value   bool
	}

	var settings []SettingsStruct
	for name, data := range utils.SetMapTable {
		settings = append(settings, SettingsStruct{
			Name:    name,
			Comment: data.Comment,
			value:   utils.GetBitByName(user.Setting, name),
		})
	}

	c.JSON(200, settings)
}

func SetSettings(user *User, c *gin.Context) {

	var request struct {
		Name  string
		Value bool
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求不正确", err))
		return
	}

	user.Setting = utils.UpdateBitByName(
		user.Setting, request.Name, request.Value,
	)

	if err := DB.Updates(&user).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil))
}
