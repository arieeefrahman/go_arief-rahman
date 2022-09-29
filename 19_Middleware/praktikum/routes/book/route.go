package book

import (
	"prak-19-middleware/constants"
	b "prak-19-middleware/controller/book"
	m "prak-19-middleware/middleware"

	"github.com/labstack/echo/v4"
	midlib "github.com/labstack/echo/v4/middleware"
)

func SetupBookRoute(e *echo.Echo) {
	e.GET("/books", b.GetAllBooksController)
	e.GET("/books/:id", b.GetBookController)
	e.POST("/books", b.CreateBookController)
	e.DELETE("/books/:id", b.DeleteBookController)
	e.PUT("/books/:id", b.UpdateBookController)

	// log
	m.LogMiddleware(e)

	// jwt authentication
	eJwt := e.Group("/jwt")
	eJwt.Use(midlib.JWT([]byte(constants.SECRET_JWT)))
	eJwt.POST("/books", b.CreateBookController)
	eJwt.DELETE("/books/:id", b.DeleteBookController)
	eJwt.PUT("/books/:id", b.UpdateBookController)
}