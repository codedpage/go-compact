package main

import (
	. "fmt"
)

func main() {

	ch := make(chan bool)
	go hello(ch)

	Println(<-ch)

}

func hello(ch chan bool) {

	ch <- true

}
