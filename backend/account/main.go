package main

import (
	"log"

	"github.com/McaxDev/Axolotland/backend/verification/rpc"
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
}
