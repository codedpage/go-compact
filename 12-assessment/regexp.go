package main

import (
	"fmt"
	"regexp"
)

func main() {
	testArray := []string{"A", "b", "c", "1", "2", "3", "4", "5", "6", "@", "~", "D", "["}

	letterRegex := regexp.MustCompile(`^[A-Za-z]$`)
	numberRegex := regexp.MustCompile(`^\d$`)
	specialRegex := regexp.MustCompile(`^[^A-Za-z\d]$`)

	var letters, numbers, special []string

	for _, char := range testArray {
		if letterRegex.MatchString(char) {
			letters = append(letters, char)
		} else if numberRegex.MatchString(char) {
			numbers = append(numbers, char)
		} else if specialRegex.MatchString(char) {
			special = append(special, char)
		}
	}
	fmt.Println(letters, numbers, special)
}
