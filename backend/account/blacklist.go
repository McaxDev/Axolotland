package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBlackList(c *gin.Context) {

	type NumAndExp struct {
		Value  string
		Expiry time.Time
	}
	var response map[string][]NumAndExp
	var data []BlackList
	if err := DB.Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("查找失败", err, nil))
		return
	}

	for _, value := range data {
		response[value.Type] = append(response[value.Type],
			NumAndExp{Value: value.Value, Expiry: value.Expiry},
		)
	}

	c.JSON(200, utils.Resp("查找成功", nil, response))
}

func SetBlackList(user *User, c *gin.Context) {

	if !user.Admin {
		c.JSON(400, utils.Resp("你不是管理员", nil, nil))
		return
	}

	var request BlackList
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	var blackList BlackList
	if err := DB.First(
		&blackList, "type = ? AND value = ?", request.Type, request.Value,
	).Error; err == gorm.ErrRecordNotFound {
		if err := DB.Create(&request).Error; err != nil {
			c.JSON(500, utils.Resp("创建黑名单记录失败", err, nil))
			return
		}
	} else if err == nil {
		if err := DB.Model(&blackList).Where(
			"id = ?", blackList.ID,
		).Updates(&request).Error; err != nil {
			c.JSON(500, utils.Resp("更新黑名单记录失败", err, nil))
			return
		}
	} else {
		c.JSON(500, utils.Resp("查找黑名单记录失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新成功", nil, nil))
}

func DeleteBlackList(user *User, c *gin.Context) {

	if !user.Admin {
		c.JSON(400, utils.Resp("你不是管理员", nil, nil))
		return
	}

	var request BlackList
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	if err := DB.First(&request).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("不存在这个黑名单记录", nil, nil))
		return
	}

	if err := DB.Delete(&request).Error; err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}
}
