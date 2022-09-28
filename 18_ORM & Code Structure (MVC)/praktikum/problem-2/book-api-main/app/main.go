package main

import (
	"alta/book-api/routes"
	customerController "alta/book-api/api/controllers/customer"
	userController "alta/book-api/api/controllers/user"
	bookController "alta/book-api/api/controllers/book"
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDatabaseConnection(config)

	// ----------------------------------------------
	// customer
	// ----------------------------------------------
	//initiate customer model
	customerModel := models.NewCustomerModel(db)

	//initiate customer controller
	newCustomerController := customerController.NewController(customerModel)

	// ----------------------------------------------
	// user
	// ----------------------------------------------
	//initiate user model
	userModel := models.NewUserModel(db)

	//initiate user controller
	newUserController := userController.NewController(userModel)
	
	// ----------------------------------------------
	// book
	// ----------------------------------------------
	//initiate book model
	bookModel := models.NewBookModel(db)

	//initiate book controller
	newBookController := bookController.NewController(bookModel)

	//create echo http
	e := echo.New()

	//register API path and controller
	routes.RegisterPathCustomer(e, newCustomerController)
	routes.RegisterPathUser(e, newUserController)
	routes.RegisterPathBook(e, newBookController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
