package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/McaxDev/Axolotland/backend/account/rpc"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func (s *RPCServer) UpdateJWT(
	c context.Context, r *rpc.JWT,
) (*rpc.JWT, error) {

	user, err := GetUser(r.JWT)
	if err != nil {
		return nil, err
	}

	newToken, err := GetJwt(user.ID)
	if err != nil {
		return nil, err
	}

	return &rpc.JWT{JWT: newToken}, err
}

func (s *RPCServer) GetUserinfo(
	c context.Context, r *rpc.JWT,
) (*rpc.Userinfo, error) {

	user, err := GetUser(r.JWT)
	if err != nil {
		return nil, err
	}

	return &rpc.Userinfo{
		Username:    user.Username,
		Avatar:      user.Avatar,
		Profile:     user.Profile,
		Admin:       user.Admin,
		Money:       int32(user.Money),
		Email:       user.Email,
		Telephone:   user.Telephone,
		BedrockName: user.BedrockName,
		JavaName:    user.JavaName,
	}, nil
}

func GetJwt(userId uint) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(24 * time.Hour).Unix(),
		}).SignedString([]byte(Config.JwtKey))
	return "Bearer " + token, err
}

func GetUser(token string) (*User, error) {

	jwtToken, err := jwt.Parse(
		token[7:], func(t *jwt.Token) (interface{}, error) {
			return []byte(Config.JwtKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("token签名格式不正确: %v\n", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("token身份信息有误")
	}

	var user User

	if err := DB.First(
		&user, "id = ?", uint(claims["userId"].(float64)),
	).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New("找不到用户")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
