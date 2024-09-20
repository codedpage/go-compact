package main

import (
	. "fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go oddEven(ch)

	for v := range ch {
		Println(v)
	}
	time.Sleep(2 * time.Second)
}

func oddEven(ch chan string) {

	for i := 10; i < 20; i++ {

		if i%2 == 0 {
			//Println("even : ", i)
			ch <- Sprintf("even %d", i)
		} else {
			//Println("odd : ", i)
			ch <- Sprintf("odd %d", i)
		}
	}

	close(ch)
}
