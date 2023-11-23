package main

import (
	"fmt"
	"time"
)

func main0() {
	ch := make(chan int)
	go gor(ch)
	ch <- 50
	time.Sleep(250 * time.Millisecond)
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 3; i++ {
		go gor(ch)
		ch <- 50 + i
	}
	time.Sleep(250 * time.Millisecond)

}

func gor(ch chan int) {
	fmt.Println(<-ch)
}

//https://go.dev/play/p/7vUGrNPxgZJ
