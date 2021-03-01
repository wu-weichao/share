package models

import (
	"fmt"
	"gorm.io/gorm"
	"share/configs"
	"share/internal/database"
)

var db *gorm.DB

func init() {
	var err error
	db, err = database.GetDB(configs.Database)
	if err != nil {
		panic(fmt.Sprintf("db init err %v\n", err))
	}
	// model init
	// table migrate
	db.AutoMigrate(&User{}, &Tag{}, &Article{})
}
