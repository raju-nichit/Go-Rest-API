package service

import (
	"go-rest-api/dtos"
	"go-rest-api/models"
)

type UserService interface {
	SignUp(*models.UserModel) (*models.UserModel, error)
	SignIn(*models.UserModel) (*dtos.UserDTO, error)
	SignOut(authToken string) error
}
