package main

import "fmt"

func main() {

	m := map[int]string{} //m := make(map[int]string)
	m[1] = "aa"
	m[2] = "bb"
	m[3] = "cc"

	getMap(m)
}

func getMap(m map[int]string) {
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

	ki := []int{}    //[]interface{}{}
	vs := []string{} //[]interface{}{}

	for k, v := range m {
		ki = append(ki, k)
		vs = append(vs, v)
	}

	fmt.Println(ki)
	fmt.Println(vs)

}

//https://go.dev/play/p/c7mTHKW5pgl
