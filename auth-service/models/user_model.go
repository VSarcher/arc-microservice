package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" gorm:"text;not null;default:null"`
	Email        string `json:"email" gorm:"text;not null;default:null"`
	Password     string `json:"password" gorm:"text;not null;default:null"`
	PhoneNumber  string `json:"phone_number" gorm:"text;not null;default:null"`
	Avatar       string `json:"avatar" gorm:"text;default:null"`
	Is_Suspended bool   `json:"is_suspended" gorm:"bool;default:false"`
	Is_Admin     bool   `json:"is_admin" gorm:"bool;default:false"`
}
