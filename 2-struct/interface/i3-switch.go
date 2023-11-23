package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func (p Person) Desc() {
	fmt.Println(p.name, p.age)
}

type Describe interface {
	Desc()
}

func findType(i interface{}) {

	fmt.Printf("Type : %T : ", i)

	switch v := i.(type) {

	case Describe:
		v.Desc()

	case string:
		fmt.Println(i.(string))

	case int:
		fmt.Println(i.(int))

	default:

		fmt.Println("undefined")
	}
}

func main() {

	findType("aaa")
	findType(10)
	findType(math.Pi)
	p := Person{"aa", 20}
	findType(p)
}

//https://go.dev/play/p/FpwrAO7gJTD
