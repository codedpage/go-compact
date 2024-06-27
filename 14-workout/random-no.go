package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan interface{})
	go randomNumber(c)
	count := 0
	for val := range c {
		fmt.Print(val, "\t")
		count++
		if count == 10 {
			count = 0
			fmt.Println()
		}
	}
}

func randomNumber(c chan interface{}) {

	//Generate a random permutation of indices 0 to 100
	indices := rand.Perm(100)
	for _, idx := range indices {
		c <- idx + 1
		time.Sleep(1 * time.Second)
	}
	close(c)
}
