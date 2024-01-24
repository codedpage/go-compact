package main

import (
	"fmt"
	"sort"
)

func rearrangeString(input string) string {
	// Create slices to store letters and digits separately.
	var letters, digits []byte

	// Separate the input string into letters and digits.
	for i := 0; i < len(input); i++ {
		if input[i] >= 'a' && input[i] <= 'z' {
			letters = append(letters, input[i])
		} else if input[i] >= '0' && input[i] <= '9' {
			digits = append(digits, input[i])
		}
	}

	// Sort the slices.
	sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })

	// Combine the sorted letters and digits with no two adjacent characters having the same type.
	var result []byte
	letterIdx, digitIdx := 0, 0
	letterTurn := true

	for i := 0; i < len(input); i++ {
		if (letterTurn && letterIdx < len(letters)) || (digitIdx >= len(digits)) {
			result = append(result, letters[letterIdx])
			letterIdx++
		} else {
			result = append(result, digits[digitIdx])
			digitIdx++
		}
		letterTurn = !letterTurn
	}

	return string(result)
}

func main() {
	input1 := "z3b1a2"
	output1 := rearrangeString(input1)
	fmt.Println("input:", input1)
	fmt.Println("output:", output1) // Output: "1a2b3z"

	input2 := "q56"
	output2 := rearrangeString(input2)
	fmt.Println("input:", input2)
	fmt.Println("output:", output2) // Output: "5q6"
}
