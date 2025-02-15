package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3) // We have three goroutines

	c := make(chan string)
	c2 := make(chan string)

	go f1("abc", c, &wg)
	go f2(c, c2, &wg)
	go f3(c2, &wg)

	wg.Wait() // Wait for all goroutines to complete
}

func f1(s string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- s
	close(c) // Close the channel after sending
}

func f2(c chan string, c2 chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	val, ok := <-c
	if ok {
		c2 <- val
	}
	close(c2) // Close the channel after sending
}

func f3(c2 chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-c2)
}
