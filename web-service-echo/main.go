package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func userHandler(c echo.Context) error {
	dataType := c.Param("data")
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	fmt.Println(username)
	fmt.Println(name)
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("username: %s, Name: %s", username, name))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
		})
	}
	return c.String(http.StatusBadRequest, "json yada string parametresi kullanılmalı")
}
func addUser(c echo.Context) error {
	// user := User{}
	// body, err := ioutil.ReadAll(c.Request().Body)
	// if err != nil {
	// 	return err
	// }
	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(user)
	// return c.String(http.StatusOK, "success")

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u)
	return c.JSON(http.StatusCreated, u)
}
func main() {
	fmt.Println("Hello world")
	server := echo.New()
	server.GET("/main", mainHandler)
	server.GET("/user/:data", userHandler)
	server.POST("/user", addUser)
	server.Start(":8000")
}
