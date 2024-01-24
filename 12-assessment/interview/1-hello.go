package main //1

import "fmt" //2

func main() { //3  
    fmt.Println("Hello World") //4
}

//https://go.dev/play/p/QUc-LJ4CaOE

package main

import (
    "fmt"
    "sort"
)

func ArrayAdditionI(arr []int) string {
    // Sort the array in ascending order
    sort.Ints(arr)

    // Get the largest number and remove it from the array
    largest := arr[len(arr)-1]
    arr = arr[:len(arr)-1]

    // Define a recursive function to generate combinations
    var canAddToLargest func(int, int) bool
    canAddToLargest = func(target, index int) bool {
        if target == 0 {
            return true
        }
        if target < 0 || index < 0 {
            return false
        }
        // Include or exclude the current number in the combination
        return canAddToLargest(target-arr[index], index-1) || canAddToLargest(target, index-1)
    }

    // Check if any combination adds up to the largest number
    if canAddToLargest(largest, len(arr)-1) {
        return "true"
    }
    return "false"
}

func main() {
    arr := []int{4, 6, 23, 10, 1, 3}
    result := ArrayAdditionI(arr)
    fmt.Println(result) // Output: true
}
=========================

package main

import (
    "fmt"
    "strings"
)

func StringScramble(str1, str2 string) string {
    // Create maps to store character counts for both strings
    countMap1 := make(map[rune]int)
    countMap2 := make(map[rune]int)

    // Populate the count maps for str1 and str2
    for _, char := range str1 {
        countMap1[char]++
    }
    for _, char := range str2 {
        countMap2[char]++
    }

    // Check if str2 characters can be formed from str1
    for char, count := range countMap2 {
        // If str2 has more of a character than str1, it's not possible
        if count > countMap1[char] {
            return "false"
        }
    }

    // If all characters in str2 can be formed from str1, return "true"
    return "true"
}

func main() {
    str1 := "rkqodlw"
    str2 := "world"
    result := StringScramble(str1, str2)
    fmt.Println(result) // Output: true
}

=================
package main

import "fmt"

func SimpleSymbols(str string) string {
    // Convert the string to a slice of runes for easy character access
    chars := []rune(str)

    // Iterate through the characters starting from the second character
    for i := 1; i < len(chars)-1; i++ {
        if isLetter(chars[i]) {
            // Check if the adjacent characters are '+'
if chars[i-1] != '+' || chars[i+1] != '+' {
                return "false"
            }
        }
    }

    // If no letters violated the rule, return "true"
    return "true"
}

func isLetter(char rune) bool {
    // Check if the character is a letter (a to z or A to Z)
return (char >= 'A' && char <= 'Z')
}

func main() {
    input := "++d+===+c++==a"
    result := SimpleSymbols(input)
    fmt.Println(result) // Output: false
}
