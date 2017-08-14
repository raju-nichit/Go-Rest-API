package webservices

import (
	//	_ "Common-Utility/template"
	"encoding/json"
	"go-rest-api/dao"
	"go-rest-api/exceptions"
	"go-rest-api/models"
	"go-rest-api/services"
	"go-rest-api/utils"
	"io/ioutil"
	"log"
	"net/http"

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
	v1 := router.Group("go-rest-api/api/v1/user")
	{
		v1.POST("/signup", signUp)
		v1.POST("/signIn", signIn)
	}
	router.Run()
}

func signUp(c *gin.Context) {
	log.Println("<-----------Sign up function called---------------> ")
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

func signIn(c *gin.Context) {
	log.Println("<--------------signIn webservice-------------->")
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
		c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "message": "User login successfully!", "object": userModel})
	}

}
