package book

import (
	b "prak-18-prob-2/controller/book"

	"github.com/labstack/echo/v4"
)

func SetupBookRoute(e *echo.Echo) {
	e.GET("/books", b.GetAllBooksController)
	e.GET("/books/:id", b.GetBookController)
	e.POST("/books", b.CreateBookController)
	e.DELETE("/books/:id", b.DeleteBookController)
	e.PUT("/books/:id", b.UpdateBookController)
}