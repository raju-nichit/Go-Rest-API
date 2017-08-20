package service

import (
	"github.com/raju-nichit/Go-Rest-API/dtos"
	"github.com/raju-nichit/Go-Rest-API/models"
)

//UserService --
type UserService interface {
	//SignUp :SignUp here
	SignUp(*models.UserModel) (*models.UserModel, error)
	//SignIn :SignIn API here
	SignIn(*models.UserModel) (*dtos.User, error)
	//SignOut :here
	SignOut(authToken string) error
}
