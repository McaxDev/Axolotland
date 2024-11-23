package main

import (
	"context"
	"fmt"
	"time"

	"github.com/McaxDev/Axolotland/backend/gameapi/rpc"
)

func (s *RPCServer) GameBind(
	c context.Context, r *rpc.BindRequest,
) (*rpc.BindResponse, error) {

	value, exists := BindCodes[r.Player]
	if !exists {
		return &rpc.BindResponse{Success: false}, fmt.Errorf(
			"请先在服务器里使用'/axo bind'，%s\n", r.Player,
		)
	}

	if time.Now().After(value.Expiry) {
		return &rpc.BindResponse{Success: false}, fmt.Errorf(
			"绑定码已过期，请重新输入'/axo bind'，%s\n", r.Player,
		)
	}

	return &rpc.BindResponse{Success: true}, nil
}
