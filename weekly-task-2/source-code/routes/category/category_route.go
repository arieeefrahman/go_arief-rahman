package category

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) {
	e.GET("/categories", controllers.GetAllCategory)
	e.GET("/categories/:id", controllers.GetCategoryById)
	e.POST("/categories", controllers.CreateCategory)
}