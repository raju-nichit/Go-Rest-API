package webservices

import (
	//	_ "Common-Utility/template"

	"Go-Rest-API/middlewares"
	"encoding/json"
	"go-rest-api/dao"
	"go-rest-api/exceptions"
	"go-rest-api/models"
	"go-rest-api/services"
	"go-rest-api/utils"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	UserService    service.UserService = &service.UserServiceImpl{}
	UserDao        dao.UserDAO         = &dao.UserDAOImpl{}
	tokenGenerator utils.TokenGenerator
)

type UserController struct {
}

func RounterConfig() {
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())
	apiv1 := router.Group("go-rest-api/api/v1/user")
	{
		apiv1.POST("/signup", signUp)
		apiv1.POST("/signIn", signIn)
		apiv1.GET("/signOut", SignOut)
	}
	router.Run(":80")
}

func signUp(c *gin.Context) {
	log.Println("<-----------Sign up function called---------------> ")
	log.Println("API Path:\t", c.Request.URL)
	payload, _ := ioutil.ReadAll(c.Request.Body)
	var requestObject models.UserModel
	err := json.Unmarshal(payload, &requestObject)
	if err != nil {
		log.Println("Error while parsing JSON", err)
	}
	log.Println(requestObject)
	UModel, _ := UserService.SignUp(&requestObject)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": UModel})
}

/**********************
*
	@@@SignIn API @@@@

***/
func signIn(c *gin.Context) {
	log.Println("<--------------signIn webservice-------------->")
	log.Println("API Path:\t", c.Request.URL)
	payload, _ := ioutil.ReadAll(c.Request.Body)
	var requestObject models.UserModel
	err := json.Unmarshal(payload, &requestObject)
	if err != nil {
		log.Println("Error while parsing JSON", err)
	}
	userModel, err := UserService.SignIn(&requestObject)
	if err != nil {
		var ErrType = err
		if ErrType == exceptions.UserServiceException("Invalid password.") {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid password.", "status": http.StatusInternalServerError})
			c.Abort()
			return
		} else if ErrType == exceptions.UserServiceException("Email id not register  with Vesica.") {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Email id not register  with Vesica.", "status": http.StatusInternalServerError})
			c.Abort()
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User login successfully!", "object": userModel})
	}
}

/**********************
*
*	@@@SignIn Out API  @@@@
*
***/

func SignOut(c *gin.Context) {
	log.Info("<------------SignOut Webservice--------------->")
	UserService.SignOut(c.Request.Header.Get("authToken"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User SingOut successfully!"})
}
