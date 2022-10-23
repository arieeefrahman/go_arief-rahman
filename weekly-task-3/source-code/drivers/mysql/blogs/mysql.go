package blogs

import (
	"simple-blog/businesses/blogs"

	"gorm.io/gorm"
)

type blogRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) blogs.Repository {
	return &blogRepository{
		conn: conn,
	}
}

func (br *blogRepository) GetAll(name string) []blogs.Domain {
	var rec []Blog
	
	if name != "" {
		br.conn.Preload("Category").Where("name = ?", name).First(&rec)
	} else {
		br.conn.Preload("Category").Find(&rec)
	}
	
	blogDomain := []blogs.Domain{}

	for _, blog := range rec {
		blogDomain = append(blogDomain, blog.ToDomain())
	}

	return blogDomain
}

func (br *blogRepository) GetByID(id string) blogs.Domain {
	var blog Blog
	br.conn.Preload("Category").First(&blog, "id = ?", id)

	return blog.ToDomain()
}

func (br *blogRepository) GetByCategoryID(categoryId string) []blogs.Domain {
	var rec []Blog
	br.conn.Preload("Category").Where("category_id = ?", categoryId).Find(&rec)
	blogDomain := []blogs.Domain{}

	for _, blog := range rec {
		blogDomain = append(blogDomain, blog.ToDomain())
	}

	return blogDomain
}

func (br *blogRepository) Create(blogDomain *blogs.Domain) blogs.Domain {
	rec := FromDomain(blogDomain)
	result := br.conn.Create(&rec)
	result.Last(&rec)

	return rec.ToDomain()
}

func (br *blogRepository) Update(id string, blogDomain *blogs.Domain) blogs.Domain {
	var blog blogs.Domain = br.GetByID(id)
	
	updatedBlog := FromDomain(&blog)
	updatedBlog.Title = blogDomain.Title
	updatedBlog.Content = blogDomain.Content
	updatedBlog.CategoryID = blogDomain.CategoryID

	br.conn.Save(&updatedBlog)

	return updatedBlog.ToDomain()
}

func (br *blogRepository) Delete(id string) bool {
	var blog blogs.Domain = br.GetByID(id)
	deletedBlog := FromDomain(&blog)
	result := br.conn.Delete(&deletedBlog)

	// if result.RowsAffected == 0 {
	// 	return false
	// }

	return result.RowsAffected != 0
}