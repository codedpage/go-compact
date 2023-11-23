package main

import "fmt"

//deadlock
func main1() {
	ch := make(chan string)
	select {
	case <-ch:
	}
}

//deadlock handling
func main() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}
