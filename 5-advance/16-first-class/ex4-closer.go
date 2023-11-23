package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {

	//5 - Closures : a anonymous functions which can access the variables defined outside of function
	a5 := 5
	func() {
		fmt.Println("a5 =", a5)
	}()

	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

//https://go.dev/play/p/134NiQGPOcS
