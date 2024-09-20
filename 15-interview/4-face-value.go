package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Map to store the counts of each card face value
	cardCounts := make(map[int]int)

	// Generate cards for even numbers in the range 1 to 20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			cardNumber := rand.Intn(13) + 1
			cardCounts[cardNumber]++
			fmt.Printf("Even number %d generates a Clubs card: %d of Clubs\n", i, cardNumber)
		}
	}

	// Compute the total face value of cards by face type
	totalFaceValue := 0
	for faceValue, count := range cardCounts {
		totalFaceValue += faceValue * count
		fmt.Printf("Face value %d occurs %d times\n", faceValue, count)
	}

	// Print the total face value
	fmt.Printf("Total face value of all generated cards: %d\n", totalFaceValue)
}
