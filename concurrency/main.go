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
	var wg sync.WaitGroup
	wg.Add(2)

	go saying("something", &wg)
	go saying("book fair", &wg)
	fmt.Println("I'm listening.")

	wg.Wait()
}
