package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
)

// 开启容器
func StartContainer(name string) error {
	return DockerClient.ContainerStart(
		context.Background(), name, container.StartOptions{},
	)
}

// 关闭容器
func StopContainer(name string) error {
	return DockerClient.ContainerStop(
		context.Background(), name, container.StopOptions{},
	)
}

// 创建文件夹
func CreateFolder(path string) error {
	if result, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	} else if !result.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

// 根据服务器代号找到服务器
func GetServerByName(name string) (*Server, error) {

	for index := range Servers {
		if temp := &Servers[index]; temp.Name == name {
			return temp, nil
		}
	}

	return nil, fmt.Errorf("没有这个服务器: %s\v", name)
}
