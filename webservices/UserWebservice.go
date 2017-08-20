package webservices

import (
	//	_ "Common-Utility/template"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/raju-nichit/Go-Rest-API/exceptions"
	"github.com/raju-nichit/Go-Rest-API/models"
	service "github.com/raju-nichit/Go-Rest-API/services"
	"github.com/raju-nichit/Go-Rest-API/utils"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//UserWevService --
type UserWevService struct {
}

var (
	//UserService --
	UserService    service.UserService = &service.UserServiceImpl{}
	tokenGenerator utils.TokenGenerator
)

//SignUp --
func (userWevService *UserWevService) SignUp(c *gin.Context) {
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

// SignIn --
func (userWevService *UserWevService) SignIn(c *gin.Context) {
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

/*SignOut -- *********************
*
*	@@@SignIn Out API  @@@@
*
***/
func (userWevService *UserWevService) SignOut(c *gin.Context) {
	log.Info("<------------SignOut Webservice--------------->")
	UserService.SignOut(c.Request.Header.Get("authToken"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User SingOut successfully!"})
}
