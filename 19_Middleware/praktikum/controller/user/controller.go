package user

import (
	"net/http"
	"prak-19-middleware/models"
	"prak-19-middleware/database"
	usermw "prak-19-middleware/middleware/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User
	
	if err := database.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(users) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "empty user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	var user []models.User
	
	if err := database.DB.Find(&user, "id = ?" ,id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(user) == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	
	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}

	if err := database.DB.Find(&user, "id = ?" ,id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	if err := database.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	var users []models.User
	c.Bind(&user)

	err := database.DB.Model(&users).Where("id = ?", id).
		Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}).
		Error 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	
	err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed to login",
			"error": err.Error(),
		})
	}

	token, err := usermw.CreateUserToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed to login",
			"error": err.Error(),
		})
	}

	userResponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"user": userResponse,
	})
}