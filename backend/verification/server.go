package main

import (
	"context"
	"errors"

	"github.com/dchest/captcha"
)

type Server struct {
	UnimplementedVerificationServer
}

func (s *Server) VerifyCode(
	c context.Context, r *RpcRequest,
) (*RpcResponse, error) {

	switch r.Codetype {
	case "email":
		if value, exist := emailSent[r.Number]; exist {
			if r.Authcode == value.AuthCode {
				return &RpcResponse{Success: true}, nil
			}
		}
		return &RpcResponse{Success: false}, nil
	case "sms":
		if value, exist := smsSent[r.Number]; exist {
			if r.Authcode == value.AuthCode {
				return &RpcResponse{Success: true}, nil
			}
		}
		return &RpcResponse{Success: false}, nil
	case "captcha":
		if captcha.VerifyString(r.Number, r.Authcode) {
			return &RpcResponse{Success: true}, nil
		}
		return &RpcResponse{Success: false}, nil
	default:
		return &RpcResponse{Success: false}, errors.New("invalid codetype")
	}
}
