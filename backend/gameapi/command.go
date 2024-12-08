package main

import (
	"context"
	"io"

	"github.com/McaxDev/Axolotland/backend/gameapi/rpc"
	"github.com/docker/docker/api/types/container"
)

// 向容器的程序发送命令
func (server *Server) SendCmd(
	ctx context.Context, command string,
) (string, error) {

	execute, err := DockerClient.ContainerExecCreate(
		ctx, server.Name, container.ExecOptions{
			Cmd:          []string{command},
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
	)
	if err != nil {
		return "", err
	}

	response, err := DockerClient.ContainerExecAttach(
		ctx, execute.ID, container.ExecStartOptions{Tty: true},
	)
	if err != nil {
		return "", err
	}
	defer response.Close()

	data, err := io.ReadAll(response.Reader)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *RPCServer) SendCmd(
	c context.Context, r *rpc.CmdReq,
) (*rpc.String, error) {

	var server Server
	if err := DB.First(
		&server, "name = ?", r.Server,
	).Error; err != nil {
		return new(rpc.String), err
	}

	response, err := server.SendCmd(c, r.Cmd)
	if err != nil {
		return new(rpc.String), err
	}

	return &rpc.String{Data: response}, nil
}
