package models

import (
	"fmt"
	"gorm.io/gorm"
	"share/configs"
	"share/internal/database"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = database.GetDB(configs.Database)
	if err != nil {
		panic(fmt.Sprintf("db init err %v\n", err))
	}
	// model init
	// table migrate
	DB.AutoMigrate(&User{})
}
