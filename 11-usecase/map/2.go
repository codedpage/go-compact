package main

import "fmt"

func main() {
	// Create example map.
	m := map[int]string{
		1: "aa",
		2: "bb",
		3: "cc",
	}

	// Convert map to slice of keys.
	keys := []int{}
	for key, _ := range m {
		keys = append(keys, key)
	}

	// Convert map to slice of values.
	values := []string{}
	for _, value := range m {
		values = append(values, value)
	}

	fmt.Println(keys)
	fmt.Println(values)

	ki := []interface{}{} //[]int{}
	vs := []interface{}{} //[]string{}

	for k, v := range m {
		ki = append(ki, k)
		vs = append(vs, v)
	}

	fmt.Println(ki)
	fmt.Println(vs)

}

//https://go.dev/play/p/KxVK0mBvF1J
