package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("F - 1")
	go f("GR - 2")

	//inline function
	go func() {
		fmt.Println("GR - 3")
	}() //func calling

	go func(msg string) {
		fmt.Println(msg)
	}("GR - 4") //func calling and params passing

	a := func(msg string) {
		fmt.Println(msg)
	}
	a("GR - 5") //func calling and params passing

	fmt.Scanln()
	fmt.Println("Finish - 6")
}

//https://go.dev/play/p/FtROCWmwxbp
