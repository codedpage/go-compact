package main

import (
	"fmt"
	"time"
)

func main1() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}

// ////////////////
func main2() {
	ch := make(chan bool)
	go hello(ch)

	select {
	case <-ch:
		fmt.Println("received value")

	}
}

func hello(ch chan bool) {

	fmt.Println("Hello gr")
	ch <- true
}

// ////////////
func process(ch chan string) {
	time.Sleep(5500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return //stop after get the respose
		default:
			fmt.Println("no value received")
		}
	}

}
