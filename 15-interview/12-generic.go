// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	age    int
	salary int
}

func main() {
	fmt.Println("Hello, 世界")

	e := []Person{
		{"dk", 51, 2000},
		{"bk", 31, 1000},
		{"ck", 41, 4000},
	}
	fmt.Println(e)

	en := make(map[string]Person)
	ea := make(map[int]Person)
	es := make(map[int]Person)

	sn := []string{}
	sa := []int{}
	ss := []int{}

	for _, v := range e {
		en[v.name] = v
		sn = append(sn, v.name)
	}
	fmt.Println(en, sn)

	for _, v := range e {
		ea[v.age] = v
		sa = append(sa, v.age)
	}
	fmt.Println(ea, sa)

	for _, v := range e {
		es[v.salary] = v
		ss = append(ss, v.salary)
	}
	fmt.Println(es, ss)
	//////////////////////

	sort.Strings(sn)
	sort.Ints(sa)
	sort.Ints(ss)

	fmt.Println(en, sn)
	fmt.Println(ea, sa)
	fmt.Println(es, ss)
}
