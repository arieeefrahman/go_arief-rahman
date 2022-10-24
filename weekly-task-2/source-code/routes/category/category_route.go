package category

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) {

	categories := e.Group("/categories")
	categories.GET("", controllers.GetAllCategory).Name = "get-all-categories"
	categories.GET("/categories/:id", controllers.GetCategoryById).Name = "get-category-by-id"
	categories.POST("/categories", controllers.CreateCategory).Name = "create-category"
}