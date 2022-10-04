package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID 		uint `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Token	string `json:"token"`
}

type UserInput struct {
	ID 			uint `json:"id" validate:"required"`
	Name 		string `json:"name" validate:"required"`
	Email 		string `json:"email" validate:"required"`
	Password 	string `json:"password" validate:"required"`
}
