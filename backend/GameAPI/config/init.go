package config

import (
	"log"

	"github.com/docker/docker/client"
)

var (
	DockerClient *client.Client
)

func Init() {
	var err error
	if DockerClient, err = client.NewClientWithOpts(
		client.WithHost(DockerHost),
		client.WithAPIVersionNegotiation(),
	); err != nil {
		log.Fatalf("Failed to connect to Docker daemon: %v\n", err)
	}
}
