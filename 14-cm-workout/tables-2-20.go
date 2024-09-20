package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go table2_20(c)

	count := 0
	for val := range c {
		fmt.Print(val, "\t")
		count++
		if count == 10 {
			count = 0
			fmt.Println()
		}
	}

}

func table2_20(c chan int) {

	//indices := rand.Perm(20)
	indices := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, n := range indices {
		//n = n + 1
		if n == 1 {
			continue
		}
		for i := n; i <= n*10; i = i + n {
			c <- i
			time.Sleep(1 * time.Second)
		}
	}
	close(c)

}
