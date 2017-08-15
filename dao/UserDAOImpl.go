package dao

import (
	_ "database/sql"
	"fmt"
	"go-rest-api/config"
	"go-rest-api/dtos"

	_ "github.com/go-sql-driver/mysql"
	//	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type UserDAOImpl struct {
}

func (UDao *UserDAOImpl) SaveUser(userDTO dtos.UserDTO) (dtos.UserDTO, error) {
	log.Println("<-----------------SaveUser DAO ------------->")
	//var err error
	if config.DB == nil {
		println("Hi it is null")
	}
	err := config.DB.Save(&userDTO).Error
	fmt.Println("saveuserDTO after DAOImpl save")
	config.DB.Find(&userDTO)
	fmt.Print(userDTO.GetUserId())
	return userDTO, err
}

func (UDao *UserDAOImpl) GetUserByEmail(email string) (dtos.UserDTO, error) {
	log.Println("<--------GetUserByEmail DAO---------->")
	var userDTO dtos.UserDTO
	err := config.DB.Find(&userDTO, "email = ?", email).Error
	return userDTO, err
}

func (UDao *UserDAOImpl) UpdateAuthToken(userDTO dtos.UserDTO, authToken string) error {
	err := config.DB.Table("user_dtos").Where("user_id = ?", userDTO.UserId).Update("auth_token", authToken).Error

	if err != nil {
		log.Error(err)
	}
	return err
}

func (UDao *UserDAOImpl) GetUserByAuthToken(authToken string) (dtos.UserDTO, error) {

	log.Info("<---------------GetUserByAuthToken DAO--------------->")
	println("DAO auth token", authToken)
	userDTO := dtos.UserDTO{AuthToken: authToken}
	err := config.DB.Where(&userDTO).First(&userDTO).Error
	if err != nil {
		log.Println("error inside GetUserByAuthToken DAO")
		log.Error(err)
		return userDTO, err
	}
	println("User id", userDTO.UserId)
	return userDTO, err
}

func (UDao *UserDAOImpl) SignOut(authToken string) error {
	err := config.DB.Table("user_dtos").Where("auth_token = ?", authToken).
		Update("auth_token", nil).Error
	if err != nil {
		log.Error(err)
	}
	return err
}
