package main

import (
	"fmt"
)

type Employee struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {

	fmt.Println("Hi")

	emp := []Employee{
		{Firstname: "ak", Lastname: "ss", Age: 21},
		{Firstname: "bk", Lastname: "ss", Age: 11},
		{Firstname: "ck", Lastname: "ss", Age: 18},
		{Firstname: "dk", Lastname: "ss", Age: 31},
		{Firstname: "ek", Lastname: "ss", Age: 41},
	}

	e := show(emp)
	fmt.Println(e)
}

func show(emp []Employee) []Employee {

	//emp2 := make([]Employee, 0)
	emp2 := []Employee{}

	for _, v := range emp {
		if v.Age > 20 {
			emp2 = append(emp2, v)
		}
	}

	return emp2

}
