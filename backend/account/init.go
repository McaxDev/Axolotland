package main

import (
	"time"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
)

var ChinaTime *time.Location

type RPCServer struct {
	rpc.UnimplementedAccountServer
}

func Init() error {
	var err error

	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	return nil
}
