package main

import "fmt"

//receive only
func go_in(c chan<- int) {
	c <- 10
}

//send only
func go_out(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)
	go go_in(c)
	fmt.Println(<-c)

	///////////////
	d := make(chan int)
	go go_out(d)
	d <- 50

}

//https://go.dev/play/p/IAVzCa8A7eW
