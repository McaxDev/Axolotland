package main

import (
	"encoding/json"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/McaxDev/Axolotland/backend/utils"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DockerClient  *client.Client
	Compressor    *archiver.TarGz
	AccountClient rpc.AccountClient
)

func Init() error {
	var err error

	if DockerClient, err = client.NewClientWithOpts(
		client.WithHost(config.DockerHost),
		client.WithAPIVersionNegotiation(),
	); err != nil {
		return err
	}

	Compressor = archiver.NewTarGz()

	serversByte, err := utils.ReadFile(config.ServersPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(
		serversByte, &Servers,
	); err != nil {
		return err
	}

	AccountConnect, err := grpc.NewClient(
		config.GRPCAddr.Account,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	defer AccountConnect.Close()

	AccountClient = rpc.NewAccountClient(AccountConnect)

	return nil
}
