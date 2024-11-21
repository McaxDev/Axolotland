package main

import (
	"bytes"
	"context"
	"io"

	"github.com/docker/docker/api/types/container"
)

// 向容器的程序发送命令
func (server *Server) SendCmd(
	ctx context.Context, command string,
) (string, error) {

	// 附加到容器
	response, err := DockerClient.ContainerAttach(
		ctx, server.Name, container.AttachOptions{
			Stream: true, Stdin: true, Stdout: true, Stderr: true,
		},
	)
	if err != nil {
		return "", err
	}
	defer response.Close()

	// 发送命令
	_, err = response.Conn.Write([]byte(command))
	if err != nil {
		return "", err
	}

	// 读取响应
	var message bytes.Buffer
	_, err = io.Copy(&message, response.Reader)
	if err != nil {
		return "", err
	}
	return message.String(), nil
}
