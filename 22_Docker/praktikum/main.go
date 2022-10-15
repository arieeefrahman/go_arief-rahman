package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
	Price int `json:"price"`
}

var DB *gorm.DB

func main() {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"mariadb-server",
		"3306",
		"learn_docker",
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("error: ", err)
	}

	DB.AutoMigrate(&Item{})

	log.Println("Connected to the database")

	app := echo.New()


	app.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello",
		})
	})

	app.GET("/items", func(c echo.Context) error {
		var items []Item
		if err := DB.Find(&items).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
			"data": items,
		})
	})

	app.POST("/items", func(c echo.Context) error {
		var item Item
		if err := c.Bind(&item); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}

		if err := DB.Create(&item).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "item created!",
			"data": item,
		})
	})

	app.Logger.Fatal(app.Start(":3000"))
}