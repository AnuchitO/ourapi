package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUsers(c echo.Context) error {
	uu := []User{
		User{
			ID:       1,
			Name:     "Leanne Graham",
			Username: "Bret",
			Email:    "Sincere@april.biz",
			Phone:    "1-770-736-8031 x56442",
		},
	}
	return c.JSON(http.StatusOK, uu)
}
func main() {
	e := echo.New()

	e.GET("/users", GetUsers)

	log.Fatal(e.Start(":1323"))
}
