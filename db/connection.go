package db

import (
	"os"

	"github.com/IsraelTeo/api-jwt-go/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func Connection() error {
	var err error
	GDB, err = gorm.Open(mysql.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDB() error {
	err := GDB.AutoMigrate(&model.Role{}, &model.User{})
	if err != nil {
		return err
	}
	return nil
}
