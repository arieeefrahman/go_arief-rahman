package product

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)
func SetupRoute(e *echo.Echo) {
	
	products := e.Group("/products")
	products.GET("", controllers.GetAllProducts).Name = "get-all"
	products.GET("", controllers.GetProductByName).Name = "get-by-name"
	products.GET("/:id", controllers.GetProductById).Name = "get-by-id"
	products.POST("", controllers.CreateProduct).Name = "create-product"
	products.PUT("/:id", controllers.UpdateProduct).Name = "update-product"
	products.DELETE("/:id", controllers.DeleteProduct).Name = "delete-product"
	products.GET("/category/:category_id", controllers.GetProductByCategoryId).Name = "get-product-by-category-id"
}