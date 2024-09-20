package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	GetConfig()

	r := gin.Default()
	r.Run()
}
