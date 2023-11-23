package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Desc() {
	fmt.Println(p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Desc() {

	fmt.Println(a.state, a.country)
}

type Describe interface {
	Desc()
}

func main() {

	//non pointer
	p := Person{"aa", 20}
	p.Desc()
	(&p).Desc()

	var d1 Describe
	d1 = p //allowed
	//d1 = &p //allowed
	d1.Desc()

	//with pointer //////////////////
	a := Address{"dl", "india"}
	a.Desc()
	(&a).Desc()

	var d2 Describe
	//d2 = a //not allowed
	d2 = &a //allowed
	d2.Desc()

}
