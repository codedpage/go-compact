package main

import "fmt"

type Employee struct {
	name, city string
	salary     int
}

// method
func (e Employee) display() {
	fmt.Println(e.name, e.city, e.salary)
}

// function
func display(e Employee) {
	fmt.Println(e.name, e.city, e.salary)
}

func main() {
	emp := Employee{
		name:   "aa",
		city:   "aaa",
		salary: 2000,
	}

	display(emp)
	emp.display()

}
