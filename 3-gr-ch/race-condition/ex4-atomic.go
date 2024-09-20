package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0 // Use int64 for atomic operations

func increment(wg *sync.WaitGroup) {
	atomic.AddInt64(&x, 1) // Atomically increment x
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

/*
The issue you're facing is a race condition because multiple goroutines are modifying the shared variable x
concurrently without synchronization. One way to resolve this issue without using a mutex is to use atomic operations,
specifically sync/atomic package, which provides thread-safe operations on integers. Here's how you can do it:

*/
