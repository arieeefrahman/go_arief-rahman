package response

import (
	"simple-blog/businesses/blogs"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title 	  	 string         `json:"title"`
	Content      string         `json:"content"`
	CategoryName string         `json:"category_name"`
	CategoryID   uint 		    `json:"category_id"`
}

func FromDomain(domain blogs.Domain) Blog {
	return Blog{
		ID: 		  domain.ID,
		Title: 		  domain.Title,
		Content:      domain.Content,
		CategoryName: domain.CategoryName,
		CategoryID:   domain.CategoryID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}