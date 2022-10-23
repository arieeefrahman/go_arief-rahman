package blogs

import (
	blogUsecase "simple-blog/businesses/blogs"
	"simple-blog/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID         uint					`json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time			`json:"created_at"`
	UpdatedAt  time.Time			`json:"updated_at"`
	DeletedAt  gorm.DeletedAt		`json:"deleted_at"`
	Title 	   string				`json:"title"`
	Content    string				`json:"content"`
	Category   categories.Category 	`json:"category"`
	CategoryID uint					`json:"category_id"`
}

func FromDomain(domain *blogUsecase.Domain) *Blog {
	return &Blog{
		ID: domain.ID,
		Title: 		domain.Title,
		Content: 	domain.Content,
		CategoryID: domain.CategoryID,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
		DeletedAt:	domain.DeletedAt,
	}
}

func (rec *Blog) ToDomain() blogUsecase.Domain {
	return blogUsecase.Domain{
		ID: 			rec.ID,
		Title: 			rec.Title,
		Content: 		rec.Content,
		CategoryName: 	rec.Category.Name,
		CategoryID: 	rec.CategoryID,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
		DeletedAt: 		rec.DeletedAt,
	}
}