package main

import (
	"prak-19-middleware/database"
	"prak-19-middleware/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)


func main() {
	// initiliaze database
	database.Init()

	// create a new echo instance
	e := echo.New()
	
	// Route to handler function
	routes.InitRoutes(e)
	
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}