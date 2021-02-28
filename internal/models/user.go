package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"size:100" json:"name"`
	Password string `gorm:"size:255" json:"password"`
	Avatar   string `gorm:"size:512" json:"avatar"`
	Email    string `gorm:"size:255" json:"email"`
	Phone    string `gorm:"size:50" json:"phone"`
	Type     int    `gorm:"comment:-1 admin 1 user" json:"type"`
	Status   int    `gorm:"comment:1 normal 0 inactive -1 disable" json:"status"`
}
