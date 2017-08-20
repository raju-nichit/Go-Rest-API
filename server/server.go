package server

import (
	"github.com/gin-gonic/gin"
	"github.com/raju-nichit/Go-Rest-API/middlewares"
	"github.com/raju-nichit/Go-Rest-API/webservices"
)

//NewRouter --
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
