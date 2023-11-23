package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("@ ", id, " -- $", j)
		time.Sleep(time.Second)
		fmt.Println("@ ", id, " ============ #", j)
		results <- fmt.Sprintf("## - %d", j)
	}
}

func main() {

	fmt.Println("@ - worker, $ - Start, # - Finish, ## - Result")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for r := 1; r <= numJobs; r++ {
		fmt.Println("<<<< ", <-results)
	}
}
