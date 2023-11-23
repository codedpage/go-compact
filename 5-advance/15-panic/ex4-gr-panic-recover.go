package main

import (
	"fmt"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery() //it recovers only this function or called function but not goroutine.
	//b()              //recovered
	go b() // not recovered

	fmt.Println("Inside A")
	time.Sleep(1 * time.Second)
}

func b() {
	//defer recovery()
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
}

// Inside A
// Inside B
// panic: oh! B panicked
