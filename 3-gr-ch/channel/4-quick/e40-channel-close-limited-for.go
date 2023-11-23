package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-c)
	}
}

func prod(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}
