package main

import "fmt"

// unbuffered : syncronous

func main() {

	fmt.Println("main")

	done := make(chan bool)

	//1st call
	go hello(done)
	<-done

	//2nd call
	go hello(done)
	<-done

}

func hello(done chan bool) {

	fmt.Println("gr")
	done <- true
}
