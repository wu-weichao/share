package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Tags        string `gorm:"size:255" json:"tags"`
	Title       string `gorm:"size:255;index;" json:"title"`
	Cover       string `gorm:"size:512;" json:"cover"`
	Key         string `gorm:"size:255;" json:"key"`
	Description string `gorm:"size:512;" json:"description"`
	Content     string `gorm:"type:text;" json:"content"`
	View        int    `gorm:"default:0;" json:"view"`
	Status      int    `gorm:"default:1;comment: 1 enable 0 disable" json:"status"`
}
