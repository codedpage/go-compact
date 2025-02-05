package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")

	ch := make(chan int)

	//ch <- 10
	//ch <- 20
	//ch <- 30

	//fmt.Println(<-ch)

	go hello(ch)
	go hello1(ch)
	time.Sleep(100 * time.Millisecond)

}

func hello(ch chan int) {
	ch <- 10
}

func hello1(ch chan int) {
	time.Sleep(1 * time.Millisecond)
	fmt.Println(<-ch)
}
