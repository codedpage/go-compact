package main

import "fmt"

type name struct {
	fn  string
	age int
}

func main() {
	n1 := name{"aa", 50}
	n2 := name{"aa", 50}
	n3 := name{"bb", 20}

	if n1 == n2 {
		fmt.Println("equal")
	}
	if n1 != n3 {
		fmt.Println("not euqal")
	}
}

//https://go.dev/play/p/tv8vgPFfLUz