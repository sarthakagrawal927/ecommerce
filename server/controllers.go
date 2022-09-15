package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func testServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	defer Recovery()
	user := getUserByMail(c.FormValue("email"))
	return c.JSON(http.StatusOK, user)
}

func getUsers(c echo.Context) error {
	defer Recovery()
	users := getAllUsersFromDB()
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	defer Recovery()
	user := createUserDB(c.FormValue("name"), c.FormValue("age"), c.FormValue("email"))
	return c.JSON(http.StatusOK, user)
}

func deleteUserByEmail(c echo.Context) error {
	defer Recovery()
	deleteCount := deleteUserByEmailDB(c.FormValue("email"))
	return c.String(http.StatusOK, fmt.Sprintf("Deleted Count: %d", deleteCount))
}
