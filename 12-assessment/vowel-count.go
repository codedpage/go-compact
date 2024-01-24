package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string

	// Prompt the user //
	fmt.Print("Enter your name: ")

	// Read input into the variable
	_, err := fmt.Scanf("%s %s", &a, &b)

	// Check for errors
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Count the vowels
	vowelCount := 0
	vowels := "aeiou"

	c := a + " " + b
	c = strings.ToLower(c)

	v := make(map[string]int)

	for _, char := range c {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
			v[string(char)]++
		}
	}

	fmt.Println(vowelCount, v)
}
