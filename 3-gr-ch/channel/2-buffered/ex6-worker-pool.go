package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {

		sum := 0
		number := job.randomno
		for number != 0 {
			digit := number % 10
			number = number / 10
			sum += digit
		}

		output := Result{job, sum}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()

	noOfJobs := 10
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 3
	createWorkerPool(noOfWorkers)

	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)

	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

//https://go.dev/play/p/TAOfi6fLVyU
