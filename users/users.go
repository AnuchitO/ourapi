package users

import (
	"encoding/json"
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
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	uu := []User{}
	err = json.NewDecoder(resp.Body).Decode(&uu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, uu)
}
