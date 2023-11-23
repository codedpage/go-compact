package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	l, w int
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() int {
	return r.l * r.w
}

func (c Circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func main() {

	r := Rectangle{10, 20}
	c := Circle{7}

	fmt.Println(r.Area())
	fmt.Println(c.Area())
}

//https://go.dev/play/p/VoaR_lphfiy
