package main

import (
	. "fmt"
)

func main() {

	ch := make(chan int)
	go hello(ch)

	for {
		if v, f := <-ch; f {
			Println(v)
		}
	}

}

func hello(ch chan int) {

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}
