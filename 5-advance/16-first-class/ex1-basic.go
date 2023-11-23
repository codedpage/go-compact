package main

import (
	"fmt"
)

func main1() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
}

func main2() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

func main() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

//https://go.dev/play/p/Xv28paORuK0
