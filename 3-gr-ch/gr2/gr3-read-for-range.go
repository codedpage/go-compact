package main

import (
	. "fmt"
)

func main() {
	ch := make(chan int)
	go hello(ch)

	for j := range ch {
		Println(j)
	}
}

func hello(ch chan int) {

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}
