/*
Getting stack trace after recover
If we recover a panic, we loose the stack trace about the panic.
Even in the program above after recovery, we lost the stack trace.

There is a way to print the stack trace using the PrintStack function of the Debug package
*/

package main

import (
	"fmt"
	//"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		//debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

//https://go.dev/play/p/LYxcJUN9mg7
