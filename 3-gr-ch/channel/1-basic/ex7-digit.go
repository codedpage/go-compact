package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		number = number / 10
		sum += digit
	}

	squareop <- sum
}

func main() {
	number := 123 // 6
	ch := make(chan int)
	go calcSquares(number, ch)

	squares := <-ch
	fmt.Println("Final output", squares)
}

// //////////////////////////

func calcSquares1(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit
	}
	squareop <- sum
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
