package dao

import (
	"fmt"

	"github.com/raju-nichit/Go-Rest-API/config"
	"github.com/raju-nichit/Go-Rest-API/dtos"
	log "github.com/sirupsen/logrus"
)

//UserDAOImpl --
type UserDAOImpl struct {
}

//SaveUser -- add user
func (UDao *UserDAOImpl) SaveUser(userDTO dtos.User) (dtos.User, error) {
	log.Println("<-----------------SaveUser DAO ------------->")
	//var err error
	if config.DB == nil {
		println("Hi it is null")
	}
	err := config.DB.Save(&userDTO).Error
	fmt.Println("saveuserDTO after DAOImpl save")
	config.DB.Find(&userDTO)
	fmt.Print(userDTO.UserId)
	return userDTO, err
}

//GetUserByEmail -- chcek user record exist or not using email
func (UDao *UserDAOImpl) GetUserByEmail(email string) (dtos.User, error) {
	log.Println("<--------GetUserByEmail DAO---------->")
	var userDTO dtos.User
	err := config.DB.Find(&userDTO, "email = ?", email).Error
	return userDTO, err
}

//UpdateAuthToken -- update authToken after login and logout
func (UDao *UserDAOImpl) UpdateAuthToken(userDTO dtos.User, authToken string) error {
	err := config.DB.Table("user_dtos").Where("user_id = ?", userDTO.UserId).Update("auth_token", authToken).Error

	if err != nil {
		log.Error(err)
	}
	return err
}

//GetUserByAuthToken --
func (UDao *UserDAOImpl) GetUserByAuthToken(authToken string) (dtos.User, error) {

	log.Info("<---------------GetUserByAuthToken DAO--------------->")
	println("DAO auth token", authToken)
	userDTO := dtos.User{AuthToken: authToken}
	err := config.DB.Where(&userDTO).First(&userDTO).Error
	if err != nil {
		log.Println("error inside GetUserByAuthToken DAO")
		log.Error(err)
		return userDTO, err
	}
	println("User id", userDTO.UserId)
	return userDTO, err
}

//SignOut --
func (UDao *UserDAOImpl) SignOut(authToken string) error {
	err := config.DB.Table("user_dtos").Where("auth_token = ?", authToken).
		Update("auth_token", nil).Error
	if err != nil {
		log.Error(err)
	}
	return err
}
