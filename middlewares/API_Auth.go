package middlewares

import (
	"go-rest-api/dao"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	UserDao dao.UserDAO = &dao.UserDAOImpl{}
)

func APIInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.String(), "SignIn") || !strings.Contains(c.Request.URL.String(), "SignUp") {
			authToken := c.Request.Header.Get("authToken")
			if authToken == "" {
				c.AbortWithStatus(401)
				return
			}
			_, err := UserDao.GetUserByAuthToken(authToken)
			if err != nil {
				c.AbortWithStatus(401)
				return
			} else {
				// println(userDTO)
				println("Url intercept successfully")
				c.Next()
				return
			}
			//inctercept url using db
		}

	}
}
