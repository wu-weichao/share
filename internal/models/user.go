package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	Model

	Name     string `gorm:"size:100" json:"name"`
	Password string `gorm:"size:255" json:"-"`
	Avatar   string `gorm:"size:512" json:"avatar"`
	Email    string `gorm:"index;size:255" json:"email"`
	Phone    string `gorm:"index;size:50" json:"phone"`
	Type     int    `gorm:"comment:-1 admin 1 user" json:"type"`
	Status   int    `gorm:"comment:1 enable 0 inactive -1 disable" json:"status"`
}

const (
	UserTypeAdmin  = -1
	UserTypeNormal = 1
)

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
