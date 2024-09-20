// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func show(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {

	ch := make(chan int)
	go show(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	time.Sleep(5)
}
