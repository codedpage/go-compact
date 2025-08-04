package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan interface{})

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(19) + 2

	go table(randomNumber, c)
	//go table(12, c)

	for val := range c {
		fmt.Println(val)
	}
}

func table(n int, c chan interface{}) {

	for i := n; i <= n*10; i = i + n {
		c <- i
		time.Sleep(2 * time.Second)
	}
	close(c)
}
