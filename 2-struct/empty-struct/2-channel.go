package main

import (
	"fmt"
	"time"
)

// Used for signaling (e.g., in channels).

func worker(done chan struct{}) {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second) // Simulating work
	fmt.Println("Done!")
	done <- struct{}{} // Sending a signal
}

func main() {
	done := make(chan struct{})

	go worker(done)

	<-done // Waiting for signal
	fmt.Println("Main function exiting")
}
