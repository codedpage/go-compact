package main

import (
	"fmt"
)

func main() {

	var c string
	fmt.Print("Enter a char ")
	fmt.Scanf("%s", &c)

	switch c {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

}
