package routes

import (
	"weekly-task-2/routes/category"
	"weekly-task-2/routes/product"

	"github.com/labstack/echo/v4"
)
func InitRoutes(e *echo.Echo) {
	product.SetupRoute(e)
	category.SetupRoute(e)
}