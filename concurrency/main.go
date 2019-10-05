package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func saying(word string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("you said: %s : %d\n", word, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go func() {
		fmt.Println("hello")
		data := <-c
		fmt.Println("data in channel: ", data)
	}()

	fmt.Println("before send")
	time.Sleep(1 * time.Second)
	c <- "some data."
	fmt.Println("sent")

	time.Sleep(1 * time.Second)

}
