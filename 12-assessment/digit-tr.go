package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 232
	result := findDigit(num)
	fmt.Printf("Output: %d\n", result)

	num = 8124
	result = findDigit(num)
	fmt.Printf("Output: %d\n", result)

}

func findDigit(num int) int {
	// Convert num to string to iterate through its digits
	numStr := strconv.Itoa(num)

	// Initialize a count for digits that divide the number
	count := 0

	// Iterate through each digit and check if it divides the number
	for _, digit := range numStr {
		digitValue, _ := strconv.Atoi(string(digit))
		if digitValue != 0 && num%digitValue == 0 {
			count++
		}
	}

	return count
}

/*
Example 1:

Input: num=232
Explanation: 232 is divisible by 2, but not 3. Since 2 occurs twice as a digit, we returm 2
Output: 2

Example 2:

Input: num = 8124
Explanation: 8124 is divisible by all of its digits, hence the answesis 4

Output: 4
Constraints:

1 <= num <=  109

int does not contain 0 as one of its digits
*/
