package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 20; i++ { // Change the range as needed
		if i%2 == 0 {
			cardNumber := rand.Intn(13) + 1
			fmt.Printf("Even number %d generates a Hearts card: %d of Hearts\n", i, cardNumber)
		}
	}
}
