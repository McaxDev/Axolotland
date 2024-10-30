package main

import (
	"log"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var VerifyClient rpc.AccountClient

func main() {

	VerifyConn, err := grpc.NewClient(
		Config.VerifyAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer VerifyConn.Close()

	VerifyClient = rpc.NewAccountClient(VerifyConn)
}
