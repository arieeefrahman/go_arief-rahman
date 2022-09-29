package routes

import (
	"prak-19-middleware/routes/user"
	"prak-19-middleware/routes/book"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	user.SetupUserRoute(e)
	book.SetupBookRoute(e)
}