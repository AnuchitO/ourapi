package users

import (
	"net/http"

	"github.com/AnuchitO/ourapi/typicode"
	"github.com/labstack/echo"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Decoder interface {
	Decode(result interface{}) error
}

type usersAPI struct {
	service Decoder
}

func (u *usersAPI) getUsers(c echo.Context) error {
	uu := []User{}
	err := u.service.Decode(&uu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, uu)
}

func GetUsers(c echo.Context) error {
	api := &usersAPI{
		service: typicode.Get("/users"),
	}
	return api.getUsers(c)
}
