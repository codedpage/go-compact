package main

import (
	"fmt"
)

// high order
func simple(a func(a, b int) int) int {
	return a(2, 3)
}

// return function
func sample() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {

	//high order
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(simple(f))

	//func-in-func
	s := sample()
	fmt.Println(s(10, 20))
}
