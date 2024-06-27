package main

import "fmt"

func main() {
	fmt.Println('A') // 65

	for r := 'A'; r <= 'Z'; r++ {
		fmt.Println(r, string(r)) // 65 A
	}

}
