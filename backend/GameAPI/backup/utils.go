package backup

import (
	"context"
	"fmt"
	"os"

	"github.com/McaxDev/Axolotland/backend/GameAPI/config"
	"github.com/docker/docker/api/types/container"
)

// 开启容器
func StartContainer(name string) error {
	return config.DockerClient.ContainerStart(
		context.Background(), name, container.StartOptions{},
	)
}

// 关闭容器
func StopContainer(name string) error {
	return config.DockerClient.ContainerStop(
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
