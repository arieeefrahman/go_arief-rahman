package user

import (
	"prak-testing/constants"
	u "prak-testing/controller/user"
	m "prak-testing/middleware"

	"github.com/labstack/echo/v4"
	midlib "github.com/labstack/echo/v4/middleware"
)

func SetupUserRoute(e *echo.Echo) {
	e.GET("/users", u.GetAllUsersController)
	e.GET("/users/:id", u.GetUserController)
	e.POST("/users", u.CreateUserController)
	e.DELETE("/users/:id", u.DeleteUserController)
	e.PUT("/users/:id", u.UpdateUserController)

	e.POST("/login", u.LoginUserController)

	// log
	m.LogMiddleware(e)

	// jwt authentication
	eJwt := e.Group("/jwt")
	eJwt.Use(midlib.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", u.GetAllUsersController)
	eJwt.GET("/users/:id", u.GetUserController)
	eJwt.DELETE("/users/:id", u.DeleteUserController)
	eJwt.PUT("/users/:id", u.UpdateUserController)
}