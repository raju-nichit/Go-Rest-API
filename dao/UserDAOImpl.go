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
		log.Fatalln(err)
	}
	return err
}
