package main

import "os"

var config struct {
	HTTPPort    string
	GRPCPort    string
	DockerHost  string
	ServersPath string
	GRPCAddr    struct {
		Account string
	}
}

func LoadConfig() {
	config.HTTPPort = os.Getenv("HTTP_PORT")
	config.GRPCPort = os.Getenv("GRPC_PORT")
	config.DockerHost = os.Getenv("DOCKER_HOST")
	config.ServersPath = os.Getenv("SERVERS_PATH")
	config.GRPCAddr.Account = os.Getenv("GRPC_ADDR_ACCOUNT")
}
