package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go arithmeticOperation(c, "+")

	for val := range c {
		fmt.Println(val)
	}
}

func arithmeticOperation(ch chan string, op string) {

	for i := 0; i < 15; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)

		ch <- fmt.Sprintf("%d %s %d = ?", x, op, y)
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
