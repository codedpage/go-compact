package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan interface{})
	go alphabet(c)

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

func alphabet(c chan interface{}) {

	alpha := make(map[int]string)
	for r, i := 'A', 1; r <= 'Z'; r, i = r+1, i+1 {
		alpha[i] = string(r)
	}

	indices := rand.Perm(26)
	for _, idx := range indices {
		c <- alpha[idx+1]
		time.Sleep(1 * time.Second)
	}
	close(c)
}
