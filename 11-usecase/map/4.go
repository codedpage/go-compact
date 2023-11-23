package main

import "fmt"

func main() {

	m := map[interface{}]interface{}{}
	m[1] = "aa"
	m[2] = "bb"
	m[3] = "cc"

	getMap(m)
}

func getMap(m map[interface{}]interface{}) {
	// Convert map to slice of keys.
	keys := []interface{}{}
	for key, _ := range m {
		keys = append(keys, key)
	}

	// Convert map to slice of values.
	values := []interface{}{}
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

//https://go.dev/play/p/AHq_r-IcMYz