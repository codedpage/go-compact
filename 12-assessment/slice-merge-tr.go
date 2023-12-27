package main

import "fmt"

func mergeSortedSlices(a, b []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append the remaining elements from both slices
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

func main() {
	a := []int{1, 2, 4}
	b := []int{1, 3, 4}

	output := mergeSortedSlices(a, b)
	fmt.Println(output)
}

/*
//tr
a := []int{1,2,4}
b := []int{1,3,4}
output : []int{1,1,2,3,4,4]

*/
