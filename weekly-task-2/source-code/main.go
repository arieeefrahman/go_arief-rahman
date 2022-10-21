package main

import (
	"fmt"
	"weekly-task-2/db"
	"weekly-task-2/routes"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8000"

func main() {
	//initialize database
	db.InitDB()

	//create new echo instance
	app := echo.New()

	//route to handle http requests
	routes.InitRoutes(app)

	//port
	appPort := fmt.Sprintf(":%s", DEFAULT_PORT)

	//start server, log if it fails
	app.Logger.Fatal(app.Start(appPort))
}