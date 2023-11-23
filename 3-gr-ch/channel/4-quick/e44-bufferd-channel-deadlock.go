package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 5
	fmt.Println(<-c)

	c <- 5
	fmt.Println(<-c)

}

//https://go.dev/play/p/Qw8lLJilXLK