package main

import "fmt"

func main() {

	nums := []int{20, 10, 40, 30, 50}

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums)-1; j++ {

			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println(nums)
}
