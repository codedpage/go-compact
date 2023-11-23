package main

import . "fmt"

type Desc1 interface {
	Display1()
}

type Desc2 interface {
	Display2()
}

type Desc3 interface {
	Desc1
	Desc2
}

type Employee struct {
	name string
	age  int
}

func (e Employee) Display1() {
	Println("name :", e.name)
}

func (e Employee) Display2() {
	Println("age :", e.age)
}

func main() {
	s := Employee{"aa", 25}
	s.Display1()
	s.Display2()

	var i Desc3 = s
	i.Display1()
	i.Display2()

}
