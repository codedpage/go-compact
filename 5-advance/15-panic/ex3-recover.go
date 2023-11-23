package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ==>", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recovery()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
}

func main() {
	defer fmt.Println("deferred call in main")
	lname := "singh"
	fullName(nil, &lname)
}

// #without recovery
// deferred call in main
// panic: runtime error: first name cannot be nil

// #with recovery
// recovered from ==> runtime error: first name cannot be nil
// deferred call in main
