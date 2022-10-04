package routes

import (
	"prak-testing/routes/user"
	"prak-testing/routes/book"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	user.SetupUserRoute(e)
	book.SetupBookRoute(e)
}