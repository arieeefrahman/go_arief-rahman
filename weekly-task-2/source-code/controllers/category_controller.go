package controllers

import (
	"net/http"
	"strconv"
	"weekly-task-2/db"
	"weekly-task-2/models"

	"github.com/labstack/echo/v4"
)

//handle get all categories
func GetAllCategory(c echo.Context) error {
	var category []models.Category

	if err := db.DB.Preload("Category").Find(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "all category",
		"result": category,
	})
}

//handle get category by id
func GetCategoryById(c echo.Context) error {
	var category models.Category
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.Preload("Category").First(&category, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "category by id",
		"result": category,
	})
}

//handle creating category with unique name
func CreateCategory(c echo.Context) error {
	var category models.Category
	c.Bind(&category)

	if err := db.DB.Preload("Category").Create(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "category created",
		"result": category,
	})
}