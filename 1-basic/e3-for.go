package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("Sum =", sum)

	c := 0
	for {
		c++
		fmt.Println(c)
		if c > 5 {
			break
		}
	}
}
