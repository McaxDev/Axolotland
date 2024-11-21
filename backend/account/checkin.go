package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
)

func Checkin(user *User, c *gin.Context) {

	iterator := time.Now().Day()

	if utils.GetBitByIndex(user.Checkin, iterator) {
		c.JSON(200, utils.Resp("你今天已经签到过啦", nil))
		return
	}

	utils.UpdateBitByIndex(&user.Checkin, iterator, true)
	user.Money += 1

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(500, utils.Resp("签到失败", err))
		return
	}

	c.JSON(200, utils.Resp("签到成功", nil))
}

func GetCheckin(user *User, c *gin.Context) {

	type Data struct {
		Date   int
		Status bool
	}

	var datas []Data

	for i := 1; i <= 31; i++ {
		datas = append(datas, Data{
			Date: i, Status: utils.GetBitByIndex(user.Checkin, i),
		})
	}

	c.JSON(200, utils.Resp("查询成功", datas))
}
