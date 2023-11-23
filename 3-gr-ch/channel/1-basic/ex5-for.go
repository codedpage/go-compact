package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

//https://go.dev/play/p/dYdPGrWMyPM
