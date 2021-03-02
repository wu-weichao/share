package models

import (
	"fmt"
	"gorm.io/gorm"
	"share/configs"
	"share/internal/database"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt int            `gorm:"" json:"created_at,omitempty"`
	UpdatedAt int            `gorm:"" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = database.GetDB(configs.Database)
	if err != nil {
		panic(fmt.Sprintf("db init err %v\n", err))
	}
	// model init
	// table migrate
	db.AutoMigrate(&User{}, &Tag{}, &Article{}, &ArticleTag{})
}
