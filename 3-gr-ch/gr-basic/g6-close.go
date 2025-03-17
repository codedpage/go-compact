package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	c := make(chan int, 3)
	c <- 1
	c <- 2
	close(c)
	//c <- 3

	fmt.Println(<-c, <-c, <-c) // 1 2 0
}

/*
closing a channel does not prevent reading from it. Instead,
it signals that no more values will be sent into the channel.
You can still receive remaining values from a closed channel until it is fully drained.
*/
