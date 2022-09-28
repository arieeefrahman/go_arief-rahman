package user

import (
	u "prak-18-prob-2/controller/user"

	"github.com/labstack/echo/v4"
)

func SetupUserRoute(e *echo.Echo) {
	e.GET("/users", u.GetUsersController)
	e.GET("/users/:id", u.GetUserController)
	e.POST("/users", u.CreateUserController)
	e.DELETE("/users/:id", u.DeleteUserController)
	e.PUT("/users/:id", u.UpdateUserController)
}