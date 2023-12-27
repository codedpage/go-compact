package main

import (
	"fmt"
	"time"
)

// print in sequence
func main() {
	t := time.Now()

	for i := 1; i <= 10; i++ {
		go count(i)
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Print("Time Taken : ", time.Since(t))
}

func count(i int) {
	fmt.Println(i)

}
