package routes

import (
	"simple-blog/controllers/blogs"
	"simple-blog/controllers/categories"
	"simple-blog/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware 	echo.MiddlewareFunc
	JWTMiddleware 		middleware.JWTConfig
	AuthController 		users.AuthController
	BlogController 		blogs.BlogController
	CategoryController 	categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
	
	users := e.Group("/api/v1/users")
	users.POST("/signup", cl.AuthController.Signup)
	users.POST("/login", cl.AuthController.Login)

	blog := e.Group("/api/v1/blogs", middleware.JWTWithConfig(cl.JWTMiddleware))
	blog.GET("", cl.BlogController.GetAll) // get all and get blog by name
	blog.GET("/:id", cl.BlogController.GetByID)
	blog.POST("", cl.BlogController.Create)
	blog.PUT("/:id", cl.BlogController.Update)
	blog.DELETE("/:id", cl.BlogController.Delete)
	blog.GET("/category/:category_id", cl.BlogController.GetByCategoryID)

	category := e.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))
	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.POST("/logout", cl.AuthController.Logout)
}