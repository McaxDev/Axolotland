package main

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/McaxDev/Axolotland/backend/gameapi/rpc"
)

func (s *RPCServer) LoadJSON(
	c context.Context, r *rpc.SrvAndPath,
) (*rpc.ByteSlice, error) {

	server, err := GetServerByName(r.Server)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(
		server.Path.Server, r.Path,
	))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &rpc.ByteSlice{Data: data}, nil
}
