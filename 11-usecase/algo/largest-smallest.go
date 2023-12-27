package main

import (
	"fmt"
)

func largest() {
	nums := []int{20, 10, 50, 30, 40}
	first := nums[0]
	for _, v := range nums {
		if v > first {
			first = v
		}
	}
	fmt.Println("largest = ", first)
}

func smallest() {
	nums := []int{20, 10, 50, 30, 40}
	first := nums[0]
	for _, v := range nums {
		if v < first {
			first = v
		}
	}
	fmt.Println("smallest = ", first)
}

func main() {
	largest()
	smallest()
}
