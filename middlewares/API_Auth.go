package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func APIInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.String(), "SignIn") || !strings.Contains(c.Request.URL.String(), "SignUp") {
			authToken := c.Request.Header.Get("authToken")
			if authToken == "" {
				c.AbortWithStatus(401)
				return
			}
			//inctercept url using db
		}
		c.Next()
		return
	}
}
