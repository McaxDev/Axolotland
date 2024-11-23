package main

import (
	"log"
	"net"
	"time"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
)

var restyClient *resty.Client

func main() {

	restyClient = resty.New().SetTimeout(5 * time.Second)

	GetConfig()

	lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatalln("failed to listen: " + err.Error())
	}
	s := grpc.NewServer()
	rpc.RegisterVerificationServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalln("failed to serve: " + err.Error())
		}
	}()

	r := gin.Default()
	r.GET("/:codetype/:number", HttpHandler)
	if err := r.Run(":" + config.HttpPort); err != nil {
		log.Fatalln("failed to run http server: " + err.Error())
	}
}
