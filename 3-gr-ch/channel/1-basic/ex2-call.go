/*

data := <- a // read from channel a
a <- data // write to channel a


Sends and receives are blocking by default
Sends and receives to a channel are blocking by default. What does this mean? When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.

*/

//goroutine with sleep

/*
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello world goroutine")
}
func main() {
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}

*/

// goroutine with channel
package main

import (
	"fmt"
)

// unbuffered : syncronous
func main() {

	fmt.Println("main")

	ch := make(chan bool)
	go hello(ch)
	<-ch
}

func hello(ch chan bool) {

	fmt.Println("Hello gr")
	ch <- true
}
