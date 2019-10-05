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

type usersAPI struct {
	get func(url string) (resp *http.Response, err error)
}

func (u *usersAPI) getUsers(c echo.Context) error {
	resp, err := u.get("https://jsonplaceholder.typicode.com/users")
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

func GetUsers(c echo.Context) error {
	api := &usersAPI{
		get: http.Get,
	}
	return api.getUsers(c)
}
