package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/verification/rpc"
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

	VerifyResponse, err := VerifyClient.VerifyCode(
		context.Background(),
		&rpc.Request{
			Codetype: "email",
			Number:   request.Email,
			Authcode: request.EmailCode,
		},
	)

	if err != nil {
	}
}
