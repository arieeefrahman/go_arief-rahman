package models

import "gorm.io/gorm"

//product has foreign key "category_id" from table category
type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `gorm:"type:text"  json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	CategoryID  int    `json:"category_id"`
	Category    Category
}

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique" json:"name"`
}