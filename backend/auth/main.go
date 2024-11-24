package main

import (
	"log"
	"net"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	Init()

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
	if err := r.Run(":" + config.HttpPort); err != nil {
		log.Fatalln("failed to run http server: " + err.Error())
	}
}
