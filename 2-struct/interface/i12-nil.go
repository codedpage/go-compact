package main

import . "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		Printf("%T %v \n", d1, d1) //<nil> <nil>
	}

	d1.Describe() //panic: runtime error: invalid memory address or nil pointer dereference
}
