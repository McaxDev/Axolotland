package main

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/get/userinfo", AuthJwtMid(GetUserInfo))
	router.GET("/get/settings", AuthJwtMid(GetSettings))
	router.POST("/set/contact", AuthJwtMid(SetContact))
	router.POST("/set/username", AuthJwtMid(SetUsername))
	router.POST("/set/password", ResetPassword)
	router.POST("/set/userinfo", AuthJwtMid(SetUserInfo))
	router.POST("/signup", AuthCaptchaMid(Signup))
	router.POST("/signout", AuthCaptchaMid(AuthJwtMid(Signout)))
	router.POST("/login", AuthCaptchaMid(Login))

	return router
}
