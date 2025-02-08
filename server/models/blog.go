package models

import "gorm.io/gorm"

type Blog struct {
	Title string `json:"title" gorm:"not null"`
	Post  string `json:"post" gorm:"not null"`
	gorm.Model
}
