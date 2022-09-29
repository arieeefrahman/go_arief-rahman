package user

import (
	"prak-19-middleware/constants"
	u "prak-19-middleware/controller/user"
	m "prak-19-middleware/middleware"

	"github.com/labstack/echo/v4"
	midlib "github.com/labstack/echo/v4/middleware"
)

func SetupUserRoute(e *echo.Echo) {
	e.GET("/users", u.GetUsersController)
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
	eJwt.GET("/users", u.GetUsersController)
	eJwt.GET("/users/:id", u.GetUserController)
	eJwt.DELETE("/users/:id", u.DeleteUserController)
	eJwt.PUT("/users/:id", u.UpdateUserController)
}