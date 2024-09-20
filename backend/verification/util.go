package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomCode() string {
	result := make([]byte, 6)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func Response(message string, data any) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}
