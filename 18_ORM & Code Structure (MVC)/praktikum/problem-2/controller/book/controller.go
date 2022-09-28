package book

import (
	"net/http"
	"prak-18-prob-2/database"
	"prak-18-prob-2/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetAllBooksController(c echo.Context) error {
	var books []models.Book
	
	if err := database.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(books) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "empty books",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all books",
		"books": books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var book []models.Book

	if err := database.DB.Find(&book, "id = ?" ,id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(book) == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get book by id",
		"book": book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := database.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new book",
		"book": book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	
	if err := database.DB.Find(&book, "id = ?" ,id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	var books []models.Book
	c.Bind(&book)

	if err := database.DB.Model(&books).Where("id = ?", id).
		Updates(map[string]interface{}{"title": book.Title, "isbn": book.Isbn, "description": book.Description}).
		Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update book",
	})
}