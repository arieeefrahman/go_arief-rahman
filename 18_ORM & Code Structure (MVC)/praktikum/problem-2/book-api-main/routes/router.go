package routes
import (
	"alta/book-api/api/controllers/customer"
	"alta/book-api/api/controllers/user"
	"alta/book-api/api/controllers/book"

	echo "github.com/labstack/echo/v4"
)

func RegisterPathCustomer(e *echo.Echo, customerController *customer.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/customers/register", customerController.PostCustomerController)
	e.POST("/customers/login", customerController.DeleteCustomerController)

	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/customers", customerController.GetAllCustomerController)
	e.GET("/customers/:id", customerController.GetCustomerController)
	e.PUT("/customers/:id", customerController.EditCustomerController)
	e.DELETE("/customers/:id", customerController.DeleteCustomerController)
}

func RegisterPathUser(e *echo.Echo, userController *user.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/users/register", userController.PostUserController)
	e.POST("/users/login", userController.DeleteUserController)

	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/users", userController.GetAllUserController)
	e.GET("/users/:id", userController.GetUserController)
	e.PUT("/users/:id", userController.EditUserController)
	e.DELETE("/users/:id", userController.DeleteUserController)
}

func RegisterPathBook(e *echo.Echo, bookController *book.Controller) {
	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/books", bookController.GetAllBookController)
	e.GET("/books/:id", bookController.GetBookController)
	e.POST("/books/register", bookController.PostBookController)
	e.PUT("/books/:id", bookController.EditBookController)
	e.DELETE("/books/:id", bookController.DeleteBookController)
}