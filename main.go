package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "hi"})
	})

	log.Fatal(e.Start(":1323"))
}
