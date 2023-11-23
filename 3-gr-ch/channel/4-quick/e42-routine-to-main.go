package main

import "fmt"

func main() {
	ch := make(chan int)
	go gor(ch)
	fmt.Println(<-ch)
}

func gor(ch chan int) {
	fmt.Println("Hello")
	ch <- 5
}

//https://go.dev/play/p/niWWl7qdeSB