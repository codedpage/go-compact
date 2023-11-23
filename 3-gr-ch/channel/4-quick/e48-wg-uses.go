package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//all 10 gr must be executed
	guaranteed()

	//not sure all 10 gr must be executed
	noGuarantee()
}

func guaranteed() {
	fmt.Println("using wg")

	t := time.Now()
	n := 5
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("Time Taken : ", time.Since(t))
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	time.Sleep(2 * time.Second)
	wg.Done()
}

// ////////////////////////////////////////////////
func noGuarantee() {
	fmt.Println("\nnot using wg")

	t := time.Now()
	n := 5

	for i := 1; i <= n; i++ {
		go process1(i)
	}
	fmt.Println("Time Taken : ", time.Since(t))
}

func process1(i int) {
	fmt.Println(i)
	time.Sleep(2 * time.Second)
}
