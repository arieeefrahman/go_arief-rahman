package product

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)
func SetupRoute(e *echo.Echo) {
	
	//route get-all-product dan get-product-by-name
	e.GET("/products", controllers.GetAllProducts)

	//route get-product-by-id
	e.GET("/products/:id", controllers.GetProductById)
	
	//route create-product
	e.POST("/products", controllers.CreateProduct)

	//route update-product
	e.PUT("/products/:id", controllers.UpdateProduct)
	
	//route delete-product
	e.DELETE("/products/:id", controllers.DeleteProduct)
	
	//route  get-product-by-category-id
	e.GET("/products/category/:category_id", controllers.GetProductByCategoryId)
}