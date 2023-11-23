package main

import "fmt"

func main() {

	number := 125

	sum := 0
	for number != 0 {
		digit := number % 10
		number = number / 10
		sum = sum + digit
	}
	fmt.Println(sum)
}
