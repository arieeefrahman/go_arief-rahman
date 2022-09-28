package routes

import (
	"prak-18-prob-2/routes/user"
	"prak-18-prob-2/routes/book"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	user.SetupUserRoute(e)
	book.SetupBookRoute(e)
}