package product

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)
func SetupRoute(e *echo.Echo) {
	e.GET("/products", controllers.GetAllProducts)
	e.GET("/products/:id", controllers.GetProductById)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)
	e.GET("/products/category/:category_id", controllers.GetProductByCategoryId)
}