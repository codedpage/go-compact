package main

import (
	"fmt"
)

func main() {
	nums := []int{20, 10, 50, 30, 40}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("largest = ", max)
}

func main2() {
	nums := []int{20, 10, 50, 30, 40}
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	fmt.Println("smallest = ", min)
}
