/* #Check if a number is Palindrome or Not
Problem Statement:
Given an integer N, return true if it is a palindrome else return false.
A palindrome is a number that reads the same backward as forward. For example, 121, 1331, and 4554 are palindromes because they remain the same when their digits are reversed.

Example 1:
Input:N = 4554
Output:Palindrome Number
Explanation: The reverse of 4554 is 4554 and therefore it is palindrome number

Example 2:
Input:N = 7789
Output: Not Palindrome
Explanation: The reverse of number 7789 is 9877 and therefore it is not palindrome
has context menu


has context menu */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	n := 1221
	fmt.Println(n)

	a := []int{}
	b := []int{}
	i := 0

	p := 0
	for n > 0 {
		m := n / 10
		r := n - m*10
		a = append(a, r)
		b = append(append([]int{}, r), b...)
		n = m
		i++
	}

	fmt.Println(a, b, p)

	// q := 0
	// for k, v := range a {
	// 	x := math.Pow10(k)
	// 	q = q + v*int(x)
	// }
	// fmt.Println(q)

	// for i := len(a), j:=0; i >= 0, j>=0; i--, j++ {
	// 	x := math.Pow10(k)
	// 	q = q + v*int(x)
	// }
	pFlag := true
	for k := range a {
		if a[k] != b[k] {
			pFlag = false
		}
	}

	if pFlag {
		fmt.Println(n, "P")
	} else {
		fmt.Println(n, "NP")
	}
}
