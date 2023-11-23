package main

import "fmt"

func main() {

	//asyncronous
	c1 := make(chan string, 2)
	c1 <- "aaa"
	c1 <- "bbb"

	fmt.Println(c1, "==", <-c1)
	fmt.Println(c1, "==", <-c1)

}

//https://go.dev/play/p/pRUHkfcY_rJ
