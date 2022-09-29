package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title 		string `json:"title" form:"title"`
	Isbn 		string `json:"isbn" form:"isbn"`
	Description string `json:"description" form:"description"`
}

type BookResponse struct {
	ID 			uint `json:"id" form:"id"`
	Title 		string `json:"title" form:"title"`
	Isbn 		string `json:"isbn" form:"isbn"`
	Description string `json:"description" form:"description"`
}