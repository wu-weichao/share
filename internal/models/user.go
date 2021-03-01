package models

import (
	"crypto/sha256"
	"encoding/hex"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"size:100" json:"name"`
	Password string `gorm:"size:255" json:"password"`
	Avatar   string `gorm:"size:512" json:"avatar"`
	Email    string `gorm:"index;size:255" json:"email"`
	Phone    string `gorm:"index;size:50" json:"phone"`
	Type     int    `gorm:"comment:-1 admin 1 user" json:"type"`
	Status   int    `gorm:"comment:1 enable 0 inactive -1 disable" json:"status"`
}

func UserGetByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserGetById(id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserEncodePassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return string(hex.EncodeToString(h.Sum(nil)))
}
