package main

import . "fmt"

func main() {
	Printf("Hello, World!")

	a := []int{1, 2, 3, 4, 5, 2, 4, 4, 6, 7}
	Println(a)
	b := map[int]int{}

	for _, e := range a {
		b[e]++
	}

	duplicates := []int{}
	for num, count := range b {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}

	Println(duplicates)
}
