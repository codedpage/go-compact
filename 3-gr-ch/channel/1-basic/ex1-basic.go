package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T %v\n\n", a, a)
	}

	var b chan int
	b = make(chan int)
	fmt.Printf("%T %v\n", b, b)
}

//https://go.dev/play/p/bDZzqcqWS_7
