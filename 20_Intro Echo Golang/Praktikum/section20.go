package main

import (
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

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user",
				"user":    user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for i, user := range users {
		if user.Id == id {
			// remove the user from the slice
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for i, user := range users {
		if user.Id == id {
			// binding new data
			updatedUser := User{}
			c.Bind(&updatedUser)

			// update the user
			updatedUser.Id = user.Id
			users[i] = updatedUser

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user":    updatedUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
