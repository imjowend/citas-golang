package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
