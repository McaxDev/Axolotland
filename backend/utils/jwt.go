package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetUserInfo(c *gin.Context, grpcAddr string) error {

	jwt := c.Request.Header.Get("Authorization")[7:]

	conn, err := grpc.Dial(
		grpcAddr, grpc.WithInsecure(), grpc.WithBlock(),
	)
	if err != nil {

	}
}
