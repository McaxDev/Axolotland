package utils

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context, grpcAddr string) {

	jwt := c.Request.Header.Get("Authorization")[7:]

}
