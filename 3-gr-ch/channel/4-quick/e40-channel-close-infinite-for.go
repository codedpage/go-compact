package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for {
		v, ok := <-c
		if !ok {
			break
		}

		fmt.Println(v, "==", ok)
	}
}

func prod(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}
