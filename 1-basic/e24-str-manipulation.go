package main

import (
	"fmt"
	. "fmt"
	"strings"
)

// concate
func main1() {

	t := fmt.Sprintf("/usr/local/%s", "help.txt")
	fmt.Println(t)

	s := "aa" + " " + "bb"
	fmt.Println(s)
}

// char occurance
func main2() {
	str := "hello"
	output := make(map[string]int)

	for k, v := range str {
		Println(k, v)
		output[string(v)]++
	}
	Println(output)
}

// string occurance
func main() {
	str := "hello baby hello kids hello"
	s := strings.Split(str, " ")
	output := make(map[string]int)

	for k, v := range s {
		Println(k, v)
		output[string(v)]++
	}
	Println(output)
}
