package main

import (
	"log"

	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var VerifyClient rpc.VerificationClient

func main() {

	VerifyConn, err := grpc.NewClient(
		Config.VerifyAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer VerifyConn.Close()

	VerifyClient = rpc.NewVerificationClient(VerifyConn)

	router := gin.Default()

	router.GET("/get/userinfo", AuthJwtMiddle(GetUserInfo))
	router.GET("/get/settings", AuthJwtMiddle(GetSettings))
	router.POST("/set/contact", AuthJwtMiddle(SetContact))
	router.POST("/set/username", AuthJwtMiddle(SetUsername))
	router.POST("/set/password", ResetPassword)
	router.POST("/set/userinfo", AuthJwtMiddle(SetUserInfo))
	router.POST("/signup", Signup)
	router.POST("/signout", AuthJwtMiddle(Signout))
	router.POST("/login", Login)

	if Config.SSL.Certificate == "" {
		err = router.Run(":" + Config.Port)
	} else {
		err = router.RunTLS(
			":"+Config.Port, Config.SSL.Certificate, Config.SSL.Private,
		)
	}

	if err != nil {
		log.Fatalf("HTTP服务器开启失败: %v\n", err)
	}
}
