package main

import (
	"fmt"
	. "fmt"
	"strconv"
)

func main() {

	echo := fmt.Println
	echo("Hello " + "World")

	fmt.Println("Hello world")
	Println("Hello world")

	//string to int
	a := "5"
	b, _ := strconv.Atoi(a)
	c := 5 + b
	fmt.Println(c)

	//int to string
	x := 20
	y := strconv.Itoa(x)
	z := "x = " + y
	fmt.Println(z)
}

//https://go.dev/play/p/oj7iN3bxnaY
