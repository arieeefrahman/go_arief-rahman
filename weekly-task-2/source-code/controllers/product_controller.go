package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"weekly-task-2/db"
	"weekly-task-2/models"

	"github.com/labstack/echo/v4"
)

//handle get all products and get product by product name
func GetAllProducts(c echo.Context) error {
	var product []models.Product

	//processing query param
	name := c.QueryParam("keyword")
	rgx := regexp.MustCompile(`[-]`)
	queryName := rgx.ReplaceAllString(name, " ")

	//if user inputted the query param, process get product by keyword
	if queryName != "" {
		if err := db.DB.Preload("Category").Where("name = ? ", queryName).First(&product).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
	
		return c.JSON(http.StatusOK, map[string]any{
			"message": "product by name",
			"result": product,
		})
	}

	// if user did not input the query param, process to get all products
	if err := db.DB.Preload("Category").Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}	

	return c.JSON(http.StatusOK, map[string]any{
		"message": "all products",
		"result": product,
	})
}

//handle get product by product id
func GetProductById(c echo.Context) error {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.Preload("Category").First(&product, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "product by id",
		"result": product,
	})
}

//handle get product by category id
func GetProductByCategoryId(c echo.Context) error {
	var product []models.Product
	categoryId, _ := strconv.Atoi(c.Param("category_id"))

	if err := db.DB.Preload("Category").Where("category_id = ? ", categoryId).Find(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "product by category id",
		"result": product,
	})
}

//handle create product
func CreateProduct(c echo.Context) error {
	var product models.Product
	c.Bind(&product)

	var categoryId = product.CategoryID

	findCategory := db.DB.First(&models.Category{ID: categoryId})

	if findCategory.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "category not found",
		})
	}

	if err := db.DB.Create(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "product created",
	})
}

//handle updating product data by id
func UpdateProduct(c echo.Context) error {
	var product models.Product
	var products []models.Product
	c.Bind(&product)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.Preload("Category").First(&products, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := db.DB.Model(&products).Where("id = ?", id).Updates(
		models.Product{
			Name: product.Name,
			Description: product.Description,
			Stock: product.Stock,
			Price: product.Price,
			CategoryID: product.CategoryID,
	}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "product updated",
	})
}

//handle deleting product by id
func DeleteProduct(c echo.Context) error {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.Preload("Category").First(&product, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := db.DB.Delete(&product, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item deleted",
	})
}