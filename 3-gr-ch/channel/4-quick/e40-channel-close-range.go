package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for v := range c {
		fmt.Println(v)
	}
}

func prod(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}
