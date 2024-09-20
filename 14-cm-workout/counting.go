package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{})
	go counting(c)
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

func counting(c chan interface{}) {

	//Generate a random permutation of indices 0 to 100

	for i := 0; i < 100; i++ {
		c <- i + 1
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
