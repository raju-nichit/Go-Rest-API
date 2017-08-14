package config

import (
	_ "database/sql"
	"go-rest-api/dtos"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB
var logger = logrus.New()

func DBConfig() {
	var err error
	logrus.Info("<--------Inside Database configuration-------------->")
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_rest?charset=utf8&parseTime=True")
	if err != nil {
		logger.Error(err)
	}
	logrus.Info("<--------Database configuration loaded-------------->")
	DB.AutoMigrate(&dtos.UserDTO{})
}
