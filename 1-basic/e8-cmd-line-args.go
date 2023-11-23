package main

import (
	"fmt"
	"os"
)

func main() {
	for a, b := range os.Args {
		fmt.Println(a, "---", b)
	}

	//fmt.Println(os.Args)
	//fmt.Printf("%T", os.Args)
}
