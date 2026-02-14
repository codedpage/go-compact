package main

import (
	"fmt"
	. "fmt"
	"os"
	"strconv"
)

func main() {

	for a, b := range os.Args {
		Println(a, b)
	}

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
