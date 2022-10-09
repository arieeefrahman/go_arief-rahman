package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type userController struct {
	useCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) *userController {
	return &userController{
		userUseCase,
	}
}

func (cr *userController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := model.User{}
		// err := db.Find(&users).Error
		// users, err := userController.useCase.UserUseCase.GetAll()
		users, err := usecase.UserUseCase.GetAll(cr.useCase)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func (cr *userController) CreateUser() echo.HandlerFunc {
	user := dto.CreateUser{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		err := usecase.UserUseCase.CreateUser(cr.useCase, user)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
