package routers

import (
	"github.com/AnuchitO/ourapi/users"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", users.GetUsers)

	return e
}
