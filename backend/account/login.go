package main

import (
	"context"
	"errors"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var request struct {
		Account  string
		Password string
		Authcode string
	}

	if err := utils.LoadBodyFromCtx(c, &request); err != nil {
		c.JSON(400, utils.Resp("请求参数有误", err, nil))
		return
	}

	var user User
	var accountType string
	if isPhone(request.Account) {
		user.Phone = request.Account
		accountType = "phone"
	} else if isEmail(request.Account) {
		user.Email = request.Account
		accountType = "email"
	} else {
		user.Name = request.Account
	}

	if request.Authcode != "" {

		if accountType == "" {
			c.JSON(400, utils.Resp("请提供有效的号码", nil, nil))
			return
		}

		response, err := AuthClient.Auth(
			context.Background(), &rpc.Authcode{
				Codetype: accountType,
				Number:   request.Account,
				Authcode: request.Authcode,
			},
		)
		if err != nil || !response.Data {
			c.JSON(400, utils.Resp("验证码验证失败", err, nil))
			return
		}

		if err := DB.First(&user).Error; errors.Is(
			err, gorm.ErrRecordNotFound,
		) {
			if err := DB.Create(&user).Error; err != nil {
				c.JSON(500, utils.Resp("新用户创建失败", err, nil))
				return
			}
		} else if err != nil {
			c.JSON(500, utils.Resp("用户查询失败", err, nil))
			return
		}
	} else if request.Password != "" {

		if err := DB.First(&user).Error; errors.Is(
			err, gorm.ErrRecordNotFound,
		) {
			c.JSON(400, utils.Resp("你尚未注册", nil, nil))
			return
		} else if err != nil {
			c.JSON(500, utils.Resp("用户查询失败", err, nil))
			return
		}

		if request.Password != user.Password {
			c.JSON(400, utils.Resp("密码不正确", nil, nil))
			return
		}
	} else {
		c.JSON(400, utils.Resp("请至少提供一种验证方式", nil, nil))
		return
	}

	token, err := GetJwt(user.ID)
	if err != nil {
		c.JSON(500, utils.Resp("用户凭证生成失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("登录成功", nil, token))
}
