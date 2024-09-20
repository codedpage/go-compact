package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan interface{})
	go pronoun(c)

	for val := range c {
		fmt.Println(val)
	}
}

func pronoun(c chan interface{}) {

	pronouns := []string{"I", "we", "you", "he", "she", "it", "this", "they", "these", "those"}

	indices := rand.Perm(len(pronouns))
	for _, idx := range indices {
		c <- pronouns[idx]
		time.Sleep(2 * time.Second)
	}
	close(c)
}
