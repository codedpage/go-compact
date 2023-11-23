package main

import (
	. "fmt"
)

func main() {
	ch := make(chan int, 5)
	ch1 := make(chan int, 5)

	go hello(ch, ch1)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	for j := range ch1 {
		Println(j)
	}

}

func hello(ch chan int, ch1 chan int) {

	for i := 1; i <= 5; i++ {

		a := <-ch
		b := a + 10
		ch1 <- b
	}
	close(ch1)

}
