package main

import (
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var request struct {
		Username     string
		Password     string
		Email        string
		EmailCode    string
		CaptchaID    string
		CaptchaValue string
	}

}
