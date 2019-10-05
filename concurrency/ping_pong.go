package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++

		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Microsecond)

		table <- ball
	}
}

func main() {

}
