package main

import (
	. "fmt"
	"time"
)

func main() {

	totalJobs := 10
	ch1 := make(chan int, totalJobs)
	ch2 := make(chan int, totalJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, ch1, ch2)
	}

	for i := 1; i <= totalJobs; i++ {
		ch1 <- i
	}
	close(ch1)

	for a := 1; a <= totalJobs; a++ {
		Println("........................................Job Done :", a, <-ch2)
	}

}

func worker(i int, ch1 chan int, ch2 chan int) {

	for j := range ch1 {
		a := j
		b := a * a
		ch2 <- b

		Println("Worker:", i, " Job:", j, " Result:", b)
		time.Sleep(1000 * time.Millisecond)
	}
	close(ch2)
}
