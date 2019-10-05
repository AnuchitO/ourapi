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
		time.Sleep(1 * time.Second)

		table <- ball
	}
}

func main() {
	table := make(chan *Ball)

	go player("fon", table)
	go player("se", table)

	table <- &Ball{}

	time.Sleep(3 * time.Second)
	<-table
}
