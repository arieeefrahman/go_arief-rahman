package request

import (
	"simple-blog/businesses/blogs"

	"github.com/go-playground/validator/v10"
)

type Blog struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

func (req *Blog) ToDomain() *blogs.Domain {
	return &blogs.Domain{
		Title: req.Content,
		Content: req.Content,
		CategoryID: req.CategoryID,
	}
}

func (req *Blog) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)

	return err
}