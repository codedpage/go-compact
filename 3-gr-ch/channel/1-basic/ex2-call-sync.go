// goroutine with channel
package main

import (
	"fmt"
)

// buffered : asyncronous => BA
// unbuffered : syncronous => US
func main() {

	fmt.Println("main")

	ch := make(chan int)
	go hello(ch)
	a, b := <-ch

	fmt.Println(a, b) //5 true
}

func hello(ch chan int) {

	fmt.Println("Hello gr")
	ch <- 5
}
