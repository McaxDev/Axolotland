package main

import (
	"context"
	"errors"

	"github.com/McaxDev/Axolotland/backend/verification/rpc"
	"github.com/dchest/captcha"
)

type Server struct {
	rpc.UnimplementedVerificationServer
}

func (s *Server) VerifyCode(
	c context.Context, r *rpc.Request,
) (*rpc.Response, error) {

	switch r.Codetype {
	case "email":
		if value, exist := emailSent[r.Number]; exist {
			if r.Authcode == value.AuthCode {
				return &rpc.Response{Success: true}, nil
			}
		}
		return &rpc.Response{Success: false}, nil
	case "sms":
		if value, exist := smsSent[r.Number]; exist {
			if r.Authcode == value.AuthCode {
				return &rpc.Response{Success: true}, nil
			}
		}
		return &rpc.Response{Success: false}, nil
	case "captcha":
		if captcha.VerifyString(r.Number, r.Authcode) {
			return &rpc.Response{Success: true}, nil
		}
		return &rpc.Response{Success: false}, nil
	default:
		return &rpc.Response{Success: false}, errors.New("invalid codetype")
	}
}
