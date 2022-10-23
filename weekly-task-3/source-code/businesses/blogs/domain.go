package blogs

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title        string
	Content      string
	CategoryName string
	CategoryID   uint
}

type Usecase interface {
	GetAll(name string) []Domain
	GetByID(id string) Domain
	GetByCategoryID(categoryId string) []Domain
	Create(blogDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll(name string) []Domain
	GetByID(id string) Domain
	GetByCategoryID(categoryId string) []Domain
	Create(blogDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}