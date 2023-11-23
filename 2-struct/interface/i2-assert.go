package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	str := "Hello World"
	describe(str) //Type = string, value = Hello World

	i := 55
	describe(i) //Type = int, value = 55

	strt := struct {
		name string
	}{
		name: "Ram",
	}
	describe(strt) //Type = struct { name string }, value = {Naveen R}

	var n interface{} = 5
	assert(n) //5 true

	var s interface{} = "abc"
	assert(s) //0 false
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
