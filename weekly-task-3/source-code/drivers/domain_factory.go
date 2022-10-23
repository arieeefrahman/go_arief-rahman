package drivers

import (
	blogDomain "simple-blog/businesses/blogs"
	categoryDomain "simple-blog/businesses/categories"
	userDomain "simple-blog/businesses/users"
	
	blogDB "simple-blog/drivers/mysql/blogs"
	categoryDB "simple-blog/drivers/mysql/categories"
	userDB "simple-blog/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewBlogRepository(conn *gorm.DB) blogDomain.Repository {
	return blogDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}