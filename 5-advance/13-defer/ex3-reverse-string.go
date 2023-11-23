package main

import (
	"fmt"
)

func main() {
	name := "welcome"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")

	// for _, v := range name {
	// 	defer fmt.Print(string(v)) //fmt.Printf("%c", v)
	// }

	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

}

//https://go.dev/play/p/PbTLnSDzueg
