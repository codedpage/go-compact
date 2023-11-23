package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	fmt.Println("Hi")

	for i := 1; i <= 5; i++ {
		go hello1(i, done)
	}

	for i := 1; i <= 5; i++ {

		fmt.Println("Finish ----", <-done)
	}

	fmt.Println("main")
}

// to maintain the order of execution
func hello(i int, done chan int) {

	var j time.Duration = time.Duration(i)
	time.Sleep(j * 250 * time.Millisecond)
	fmt.Println("Job Start ----", i, "---", j)
	done <- i
}

// no order
func hello1(i int, done chan int) {

	fmt.Println("Start ----", i)
	done <- i
}

//https://go.dev/play/p/jg957EaiwzP
