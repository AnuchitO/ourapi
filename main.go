package main

import (
	"log"

	"github.com/AnuchitO/ourapi/routers"
)

func main() {
	e := routers.NewRouter()
	log.Fatal(e.Start(":1323"))
}
