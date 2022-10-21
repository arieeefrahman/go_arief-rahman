package category

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) {

	//route get-all-categories
	e.GET("/categories", controllers.GetAllCategory)
	
	//route get-category-by-id
	e.GET("/categories/:id", controllers.GetCategoryById)

	//route create-category (unique name)
	e.POST("/categories", controllers.CreateCategory)
}