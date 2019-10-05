package main

import (
	"fmt"
	"math/rand"
	"time"
)

func saying(word string) <-chan string { // receive-only
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("you said: %s : %d", word, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	c := saying("book fair")

	for i := 0; i < 5; i++ {
		fmt.Println("You say:", <-c)
	}

	fmt.Println("Okay.")
}
