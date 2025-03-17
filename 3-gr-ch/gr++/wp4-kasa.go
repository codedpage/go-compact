package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//fmt.Println("Hello, 世界")

	var wg sync.WaitGroup
	TotalJobs := 30
	//Worker := 2

	job := make(chan int, TotalJobs)
	result := make(chan int, TotalJobs)

	wg.Add(1)
	go worker1(job, result, &wg)

	wg.Add(1)
	go worker2(job, result, &wg)

	for i := 1; i <= TotalJobs; i++ {
		job <- i
	}
	close(job)

	///////////
	wg.Wait()
	/////////////

	// //setA
	// close(result) //closing a channel does not prevent reading from it. but write protected
	// for a := range result {
	// 	fmt.Println(a)
	// }

	//setB
	for a := 1; a <= TotalJobs; a++ {
		fmt.Println(<-result)
	}

}

func worker1(j chan int, r chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for a := range j {
		time.Sleep(10 * time.Millisecond)
		r <- a + 10
	}

}
func worker2(j chan int, r chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for a := range j {
		time.Sleep(20 * time.Millisecond)
		r <- a + 100
	}

}
