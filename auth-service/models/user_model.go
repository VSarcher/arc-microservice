package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"text;not null;default:null"`
	Email    string `json:"email" gorm:"text;not null;default:null"`
	Password string `json:"password" gorm:"text;not null;default;null"`
}
