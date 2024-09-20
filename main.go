package main

import (
	. "fmt"
	"os"
)

func main() {

	for a, b := range os.Args {
		Println(a, b)
	}
}
