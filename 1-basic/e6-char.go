package main

import "fmt"

func main() {
	fmt.Println('A') // 65

	for r := 'A'; r <= 'z'; r++ {
		fmt.Println(r, string(r)) // 65 A
	}
	fmt.Printf("%c\n", 8377)
}
