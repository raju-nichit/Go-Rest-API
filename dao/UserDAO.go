package dao

import (
	"go-rest-api/dtos"
)

type UserDAO interface {
	SaveUser(dtos.UserDTO) (dtos.UserDTO, error)
	GetUserByEmail(email string) (dtos.UserDTO, error)
	UpdateAuthToken(userDTO dtos.UserDTO, authToken string) error
	GetUserByAuthToken(authToken string) (dtos.UserDTO, error)
}
