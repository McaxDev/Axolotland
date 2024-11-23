package main

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.POST("/exec/command", AuthJwtMiddle(SendCmd))
	r.POST("/exec/backup", AuthJwtMiddle(AuthAdminMiddle(WorldBackup)))
	return r
}
