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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	book := saying("book fair")
	fon := saying("Fon")

	c := fanIn(book, fon)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Okay.")
}
