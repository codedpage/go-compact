package main

import (
	"compact/1-basic/rectangle"
	"fmt"
	"log"
)

var a, b = 5, 10

func init() {

	fmt.Println("main pkg")
	if a < 10 {
		log.Println("Len < 10")
	}
	if b < 20 {
		log.Println("Width < 20")
	}
}

func main() {

	fmt.Println(float64(rectangle.Area(15, 2)))
}
