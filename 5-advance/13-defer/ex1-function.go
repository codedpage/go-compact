package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished() //2nd element : defer stack
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

// defer : ..., 3, 2, 1 element order
func main() {
	defer fmt.Println("main closed") //1st element : defer stack
	nums := []int{20, 10, 50, 30, 40}
	largest(nums)
}

// Started finding largest
// Largest number in [20 10 50 30 40] is 50
// Finished finding largest
// main closed
