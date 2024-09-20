package main

import (
	. "fmt"
)

func main() {

	totalJobs := 5
	ch1 := make(chan int, totalJobs)
	ch2 := make(chan int, totalJobs)

	go worker(ch1, ch2)
	go worker(ch1, ch2)
	go worker(ch1, ch2)

	for i := 1; i <= totalJobs; i++ {
		ch1 <- i
	}
	close(ch1)

	for a := 1; a <= totalJobs; a++ {
		Println(<-ch2)
	}

}

func worker(ch1 chan int, ch2 chan int) {

	for j := range ch1 {
		a := j
		b := a + 0
		ch2 <- b
	}
}
