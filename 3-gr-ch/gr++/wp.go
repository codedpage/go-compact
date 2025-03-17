package main

import (
	. "fmt"
)

func main() {

	job := make(chan int, 1)
	result := make(chan int, 1)
	go process(job, result)

	for i := 1; i <= 5; i++ {
		job <- i
	}

	for {
		if v, f := <-result; f {
			Println(v)
		}
	}

}

func process(job chan int, result chan int) {

	for {
		result <- 2 * <-job
	}
}
