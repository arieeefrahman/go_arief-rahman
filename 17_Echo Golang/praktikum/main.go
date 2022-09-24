package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = []User{
	{
		Id: 1, 
		Name: "Jake",
		Email: "jake123@gmail.com",
		Password: "jake8764",
	},
	{
		Id: 2,
		Name: "Diana",
		Email: "diana213@gmail.com",
		Password: "diana999876",
	},
	{
		Id: 3,
		Name: "Tony",
		Email: "tonystark@gmail.com",
		Password: "tonystarkmcu000",
	},
	{
		Id: 4,
		Name: "Wanda",
		Email: "wandamaxioff@gmail.com",
		Password: "wandamaximoff12312132131",
	},
	{
		Id: 5,
		Name: "Loki",
		Email: "lokiloki@gmail.com",
		Password: "lokisonofodin211",
	},
} 
// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	user := User{}
	id, err := strconv.Atoi(c.Param("id"))

	// check error
	if err != nil {
		log.Fatal()
	}

	checkIfIdExist := false

	for i := 0; i < len(users); i++ {
		if id == users[i].Id {
			user.Id = users[i].Id
			user.Name = users[i].Name
			user.Email = users[i].Email
			user.Password = users[i].Password
			checkIfIdExist = true
		}
	}

	// handle if user does not exist
	if !checkIfIdExist {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "failed to find user, user does not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get users by id",
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	user := User{}
	id, err := strconv.Atoi(c.Param("id"))
	
	// check error
	if err != nil {
		log.Fatal()
	}

	checkIfIdExist := false
	
	for i := len(users) - 1; i >= 0; i-- {
		if id == users[i].Id {
			user.Id = users[i].Id
			user.Name = users[i].Name
			user.Email = users[i].Email
			user.Password = users[i].Password
			
			users = append(users[:i], users[i+1:]...)
			checkIfIdExist = true
		}
	}

	// handle if user does not exist
	if !checkIfIdExist {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "failed to delete user, user does not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete users by id",
		"user": user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))

	// check error
	if err != nil {
		log.Fatal()
	}

	checkIfIdExist := false
	
	// binding data
	user := User{}
	c.Bind(&user)
	
	for i := 0; i < len(users); i++ {
		if id == users[i].Id {

			// for id update
			// prevent for updating user id
			if len(strconv.Itoa(user.Id)) > -1 {
				user.Id = id
				users[i].Id = user.Id
			}

			// for name update
			if len(user.Name) == 0 {
				// assign old value
				user.Name = users[i].Name
			} else {
				// assign new value
				users[i].Name = user.Name
			}

			// for email update
			if len(user.Email) == 0 {
				// assign old value
				user.Email = users[i].Email
			} else {
				// assign new value
				users[i].Email = user.Email
			}

			// for password update
			if len(user.Password) == 0 {
				// assign old value
				user.Password = users[i].Password
			} else {
				// assign new value
				users[i].Password = user.Password
			}

			checkIfIdExist = true
		}
	}

	// handle if user does not exist
	if !checkIfIdExist {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "failed to update data, user does not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update users by id",
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}
// ---------------------------------------------------
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
	