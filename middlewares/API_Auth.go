package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raju-nichit/Go-Rest-API/dao"
)

var (
	UserDao dao.UserDAO = &dao.UserDAOImpl{}
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("authToken")
		if !strings.Contains(c.Request.URL.String(), "SignIn") || !strings.Contains(c.Request.URL.String(), "SignUp") {
			if authToken == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":  "Invalid authToken",
					"status": http.StatusUnauthorized,
				})
				return
			} else {
				println("AuthToken:\t", authToken)
				_, err := UserDao.GetUserByAuthToken(authToken)
				if err != nil {
					println("AuthToken:\t", authToken)
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error":  "Invalid authToken",
						"status": http.StatusUnauthorized,
					})
					return
				}
			}
		}
		c.Next()
	}
}
