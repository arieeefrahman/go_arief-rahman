package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title 		string `json:"title" form:"title"`
	Isbn 		string `json:"isbn" form:"isbn"`
	Description string `json:"description" form:"description"`
}

type BookResponse struct {
	ID 			uint `json:"id"`
	Title 		string `json:"title"`
	Isbn 		string `json:"isbn"`
	Description string `json:"description"`
}

type BookInput struct {
	ID 			uint `json:"id" validate:"required"`
	Title 		string `json:"title" validate:"required"`
	Isbn 		string `json:"isbn" validate:"required"`
	Description string `json:"description" validate:"required"`
}