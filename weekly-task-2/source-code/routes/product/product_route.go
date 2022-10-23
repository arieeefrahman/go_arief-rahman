package product

import (
	"weekly-task-2/controllers"

	"github.com/labstack/echo/v4"
)
func SetupRoute(e *echo.Echo) {
	
	products := e.Group("/products")
	//route get-all-product dan get-product-by-name
	products.GET("", controllers.GetAllProducts)

	//route get-product-by-id
	products.GET("/:id", controllers.GetProductById)
	
	//route create-product
	products.POST("", controllers.CreateProduct)

	//route update-product
	products.PUT("/:id", controllers.UpdateProduct)
	
	//route delete-product
	products.DELETE("/:id", controllers.DeleteProduct)
	
	//route  get-product-by-category-id
	products.GET("/category/:category_id", controllers.GetProductByCategoryId)
}