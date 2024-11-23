package main

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/get/userinfo", AuthJwtMiddle(GetUserInfo))
	router.GET("/get/settings", AuthJwtMiddle(GetSettings))
	router.POST("/set/contact", AuthJwtMiddle(SetContact))
	router.POST("/set/username", AuthJwtMiddle(SetUsername))
	router.POST("/set/password", ResetPassword)
	router.POST("/set/userinfo", AuthJwtMiddle(SetUserInfo))
	router.POST("/signup", Signup)
	router.POST("/signout", AuthJwtMiddle(Signout))
	router.POST("/login", Login)

	return router
}
