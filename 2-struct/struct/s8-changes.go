package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (e Employee) tch(s string) {
	e.name = s
	fmt.Println("tch => ", e)
}

// pointer receiver can do permanet changes
func (e *Employee) pch(s string) {
	e.name = s
	fmt.Println("pch => ", e)
}

func main() {
	e := Employee{"aa", 20}
	fmt.Println(e)

	e.tch("bb")
	//(&e).tch("bb")
	fmt.Println(e)

	(&e).pch("cc")
	//e.pch("cc")
	fmt.Println(e)
}

func main1() {
	v := Employee{"aa", 20}
	r := &v
	fmt.Println(v)
	fmt.Println("---\n")

	//v-v
	v.tch("bb")
	fmt.Println(v)
	fmt.Println("vv---tc \n")

	//r-r ==> pch
	r.pch("cc")
	fmt.Println(v)
	fmt.Println("rr---pc\n")

	//r-v
	r.tch("dd")
	fmt.Println(v)
	fmt.Println("rv---tc \n")

	//v-r ==> pch
	v.pch("ee")
	fmt.Println(v)
	fmt.Println("vr---pc \n")
}

//https://go.dev/play/p/920r3fAohvl
