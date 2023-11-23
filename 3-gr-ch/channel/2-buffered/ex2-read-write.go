package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i + 10
		fmt.Println("ch <- ", i, " === wrote")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		//fmt.Println("read value ", v, "<- ch")
		fmt.Println("read -------------- ", v, "<- ch ")
		time.Sleep(2 * time.Second)

	}
}

//https://go.dev/play/p/ff6E54LVSCR
