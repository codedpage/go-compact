package main

import "fmt"

func main() {
	fmt.Print("Enter the no :")
	var n int
	fmt.Scanf("%d", &n)

	c1 := make(chan int)
	c2 := make(chan int)

	go sum(n, c1)
	go product(n, c2)

	a, b := <-c1, <-c2

	fmt.Println(a, "===", b)

}

func sum(n int, c1 chan int) {
	c1 <- n + 10
}

func product(n int, c2 chan int) {
	c2 <- n * 10
}

//https://go.dev/play/p/VwBJAOnXCOT
