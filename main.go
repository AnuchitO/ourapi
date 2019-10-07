package main

import (
	"log"
	"os"

	"github.com/AnuchitO/ourapi/routers"
)

func main() {
	e := routers.NewRouter()
	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}
