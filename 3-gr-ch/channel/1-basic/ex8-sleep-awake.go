package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("sleep : 4 seconds")
	time.Sleep(4 * time.Second)
	fmt.Println("awake")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}

//https://go.dev/play/p/IKSoz__50zL
