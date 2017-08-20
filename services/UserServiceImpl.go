package service

import (
	//errors --
	_ "errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/raju-nichit/Go-Rest-API/dao"
	"github.com/raju-nichit/Go-Rest-API/dtos"
	"github.com/raju-nichit/Go-Rest-API/exceptions"
	"github.com/raju-nichit/Go-Rest-API/models"
	"github.com/raju-nichit/Go-Rest-API/utils"
	log "github.com/sirupsen/logrus"
)

var (
	//UserDao --
	UserDao        dao.UserDAO = &dao.UserDAOImpl{}
	tokenGenerator utils.TokenGenerator
)

//UserServiceImpl --
type UserServiceImpl struct {
}

// SignUp API
func (UService *UserServiceImpl) SignUp(userModel *models.UserModel) (*models.UserModel, error) {
	log.Println("<-----------------Inside SaveUserInfo controller------------->")
	fmt.Println("*****In UserServiceImpl:SaveUserInfo ")
	var err error
	var checkUser dtos.User
	checkUser, _ = UserDao.GetUserByEmail(userModel.Email)
	user := dtos.User{}
	if checkUser != user {
		err = exceptions.UserServiceException("User already register with Vesica.")
		return nil, err
	}

	if userModel.Email != "" && userModel.Password != "" {
		saveUserDTO := dtos.User{}
		copier.Copy(&saveUserDTO, &userModel)
		var authToken = tokenGenerator.GenerateToken(userModel.Email)
		saveUserDTO.AuthToken = authToken
		fmt.Println("USerDTO before save in serviceImpl", saveUserDTO)
		saveUserDTO, err = UserDao.SaveUser(saveUserDTO)
		fmt.Println("USerDTO After save in serviceImpl", saveUserDTO)
		if err != nil {
			log.Print(err)
		}
		copier.Copy(&userModel, &saveUserDTO)
	} else {
		err = exceptions.UserServiceException("Email or password cannot be blank")
	}
	return userModel, err
}

//SignIn API
func (UService *UserServiceImpl) SignIn(userModel *models.UserModel) (*dtos.User, error) {
	log.Println("<-----------GetUserByEmail Service------------->")
	userDTO, err := UserDao.GetUserByEmail(userModel.Email)
	if err != nil {
		err = exceptions.UserServiceException("Email id not register  with Vesica.")
		return nil, err
	}
	if userDTO.Password == userModel.Password {
		var authToken = tokenGenerator.GenerateToken(userModel.Email)
		UserDao.UpdateAuthToken(userDTO, authToken)
		userDTO.AuthToken = authToken //SingOut --
	} else {
		err = exceptions.UserServiceException("Invalid password.")
	}
	return &userDTO, err
}

//SignOut API
func (UService *UserServiceImpl) SignOut(authToken string) error {

	log.Info("<----------SignOut service------------->")
	err := UserDao.SignOut(authToken)
	if err != nil {
		log.Error("@@@Error in Signout service@@@")
		log.Error(err)
		err = exceptions.UserServiceException("Invalid password.")
		return err
	}
	return err
}
