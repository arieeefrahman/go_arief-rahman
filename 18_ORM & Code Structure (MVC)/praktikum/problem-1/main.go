package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port string
	DB_Host string
	DB_Name string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port: "3306",
		DB_Host: "localhost",
		DB_Name: "praktikum_18",
	}
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	var err error
	// DB, err = gorm.Open("mysql", connectionString)
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
	var user []User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
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
	var user []User
	if err := DB.Delete(&user, id).Error; err != nil {
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
	user := User{}
	var users []User
	c.Bind(&user)

	// if err := DB.Model(&users).Where("id = ?", id).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	err := DB.Model(&users).Where("id = ?", id).
		Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}).
		Error 
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}