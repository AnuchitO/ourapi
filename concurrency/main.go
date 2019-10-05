package main

import (
	"fmt"
	"math/rand"
	"time"
)

func saying(word string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("you said: %s : %d", word, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go saying("book fair", c)

	for i := 0; i < 5; i++ {
		fmt.Println("You say:", <-c)
	}

	fmt.Println("Okay.")
}
