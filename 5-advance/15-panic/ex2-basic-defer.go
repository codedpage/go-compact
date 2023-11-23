package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName") //2nd element : defer stack
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
}

func main() {
	defer fmt.Println("deferred call in main") //1st element : defer stack
	lname := "singh"
	fullName(nil, &lname)
}

// deferred call in fullName
// deferred call in main
// panic: runtime error: first name cannot be nil
