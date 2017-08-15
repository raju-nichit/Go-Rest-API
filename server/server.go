package server

import (
	"Go-Rest-API/middlewares"
	"go-rest-api/webservices"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.AuthMiddleware())
	v1 := router.Group("go-rest-api/api/v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(webservices.UserWevService)
			userGroup.POST("/signup", user.SignUp)
			userGroup.POST("/signIn", user.SignIn)
			userGroup.GET("/signOut", user.SignOut)
		}
	}
	return router
}
