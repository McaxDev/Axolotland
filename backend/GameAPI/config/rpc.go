package config

import (
	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AccountClient rpc.AccountClient

func GRPCInit() error {

	AccountConnect, err := grpc.NewClient(
		GRPCAddr.Account,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	defer AccountConnect.Close()

	AccountClient = rpc.NewAccountClient(AccountConnect)

	return nil
}
