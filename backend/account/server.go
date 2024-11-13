package main

import (
	"context"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
)

type Server struct {
	rpc.UnimplementedAccountServer
}

func (s *Server) VerifyJwt(
	c context.Context, r *rpc.JwtRequest,
) (*rpc.JwtResponse, error) {

	user, err := GetUser(r.JWT)
	if err != nil {
		return nil, err
	}

	newToken, err := GetJwt(user.ID)
	if err != nil {
		return nil, err
	}

	return &rpc.JwtResponse{
		JWT:         newToken,
		Username:    user.Username,
		Avatar:      user.Avatar,
		Profile:     user.Profile,
		Admin:       user.Admin,
		Money:       int32(user.Money),
		Email:       user.Email,
		Telephone:   user.Telephone,
		BedrockName: user.BedrockName,
		JaveName:    user.JavaName,
		DstName:     user.DstName,
	}, nil
}
