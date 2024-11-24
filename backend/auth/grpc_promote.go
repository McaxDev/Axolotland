package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/auth/rpc"
)

func (s *RPCServer) Promote(
	c context.Context, r *rpc.Message,
) (*rpc.Empty, error) {
	return new(rpc.Empty), SendEmail(
		r.Number, []byte(r.Content),
	)
}
