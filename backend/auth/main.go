package main

import (
	"log"
	"net"
	"time"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	Init()

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	go func() {
		ClearSent(EmailSent, TelephoneSent, QQSent)
	}()

	lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatalln("failed to listen: " + err.Error())
	}
	s := grpc.NewServer()
	rpc.RegisterAuthServer(s, new(RPCServer))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalln("failed to serve: " + err.Error())
		}
	}()

	r := gin.Default()
	r.GET("/captcha", SendCaptcha)
	r.GET("/email/:number", SendEmailCode)
	r.GET("/telephone/:number", SendTelephone)
	r.GET("/qq/:number", SendQQCode)
	if err := r.Run(":" + config.HttpPort); err != nil {
		log.Fatalln("failed to run http server: " + err.Error())
	}
}
