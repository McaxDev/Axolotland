package utils

import "github.com/gin-gonic/gin"

func Resp(message string, data any) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}
