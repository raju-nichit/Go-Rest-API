package dao

import "github.com/raju-nichit/Go-Rest-API/dtos"

//UserDAO -- interface for user operations
type UserDAO interface {
	SaveUser(dtos.User) (dtos.User, error)
	GetUserByEmail(email string) (dtos.User, error)
	UpdateAuthToken(userDTO dtos.User, authToken string) error
	GetUserByAuthToken(authToken string) (dtos.User, error)
	SignOut(authToken string) error
}
