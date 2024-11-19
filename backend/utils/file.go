package utils

import (
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
