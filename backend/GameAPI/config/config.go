package config

import "os"

var (
	HTTPPort   string
	GRPCPort   string
	DockerHost string
	GRPCAddr   struct {
		Account string
	}
)

func LoadConfig() {
	HTTPPort = os.Getenv("HTTP_PORT")
	GRPCPort = os.Getenv("GRPC_PORT")
	DockerHost = os.Getenv("DOCKER_HOST")

	GRPCAddr.Account = os.Getenv("GRPC_ADDR_ACCOUNT")
}
