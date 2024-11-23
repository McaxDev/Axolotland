package main

import (
	"log"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthClient rpc.AuthClient

func main() {

	AuthConn, err := grpc.NewClient(
		Config.AuthAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer AuthConn.Close()

	AuthClient = rpc.NewAuthClient(AuthConn)

	router := GetRouter()

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
