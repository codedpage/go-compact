package main

import "fmt"

func main() {

	var b chan int
	b = make(chan int)
	fmt.Printf("%T %v\n", b, b)
}
