package service

import (
	_ "errors"
	"fmt"
	"go-rest-api/dao"
	"go-rest-api/dtos"
	"go-rest-api/exceptions"
	"go-rest-api/models"
	"go-rest-api/utils"

	_ "github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

var (
	UserDao        dao.UserDAO = &dao.UserDAOImpl{}
	tokenGenerator utils.TokenGenerator
)

type UserServiceImpl struct {
}

func (UService *UserServiceImpl) SignUp(userModel *models.UserModel) (*models.UserModel, error) {
	var err error
	log.Println("<-----------------Inside SaveUserInfo controller------------->")
	fmt.Println("*****In UserServiceImpl:SaveUserInfo ********")
	if userModel.GetEmail() != "" && userModel.GetPassword() != "" {
		saveUserDTO := dtos.UserDTO{
			Email:    userModel.GetEmail(),
			Password: userModel.GetPassword(),
		}
		fmt.Println("USerDTO before save in serviceImpl", saveUserDTO)
		saveUserDTO, err = UserDao.SaveUser(saveUserDTO)
		fmt.Println("USerDTO After save in serviceImpl", saveUserDTO)
		if err != nil {
			log.Print(err)
			return nil, err
		} else {
			userModel.SetUserId(saveUserDTO.GetUserId())
			userModel.SetEmail(saveUserDTO.GetEmail())
			userModel.SetAuthToken(saveUserDTO.GetAuthToken())
			return userModel, nil
		}
	} else {
		// throw  error here
		err = exceptions.UserServiceException("Email or password cannot be blank.")
		return nil, err
	}
}
func (UService *UserServiceImpl) SignIn(userModel *models.UserModel) (*dtos.UserDTO, error) {
	log.Println("<-----------GetUserByEmail Service------------->")
	userDTO, err := UserDao.GetUserByEmail(userModel.GetEmail())
	if err != nil {
		err = exceptions.UserServiceException("Email id not register  with Vesica.")
		return nil, err
	}
	if userDTO.Password == userModel.Password {
		var authToken = tokenGenerator.GenerateToken(userModel.GetEmail())
		UserDao.UpdateAuthToken(userDTO, authToken)
		userDTO.SetAuthToken(authToken)
		return &userDTO, err
	} else {
		err = exceptions.UserServiceException("Invalid password.")
		return nil, err
	}
}
