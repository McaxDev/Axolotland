package main

import "github.com/gin-gonic/gin"

func Response(message string, data any) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}
