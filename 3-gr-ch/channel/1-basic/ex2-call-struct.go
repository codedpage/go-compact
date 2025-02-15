package main

import (
	"fmt"
)

func main() {

	fmt.Println("main")

	ch := make(chan struct{})
	go hello(ch)
	<-ch
}

func hello(ch chan struct{}) {

	fmt.Println("Hello gr")
	ch <- struct{}{}
}
