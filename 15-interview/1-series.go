package main

import (
	"fmt"
)

func main() {

	n := 4

	increment := 3

	for n <= 200 {
		fmt.Println(n)

		n += increment

		increment += 2
	}
}
