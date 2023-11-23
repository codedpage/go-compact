package main

import "fmt"

type Employee struct {
	fn, ln      string
	age, salary int
}

func (a Employee) fullName() string {
	return fmt.Sprintf("Full Name: %s %s", a.fn, a.ln)
}

func main() {

	var emp Employee //zero valued structure
	fmt.Println(emp)

	e1 := Employee{"aa", "bb", 20, 5000}
	fmt.Println(e1) //{aa bb 20 5000}
	fmt.Println(e1.fullName())

	e2 := &Employee{"bb", "bbb", 30, 3000}
	fmt.Println(e2.fn, (*e2).fn) //bb bb

	e3 := &e1
	fmt.Println((*e3).fn, (*e3).ln) //aa bb

	e5 := []Employee{
		e1, *e2,
	}
	fmt.Println(e5)       //[{aa bb 20 5000} {bb bbb 30 3000}]
	fmt.Println(e5[1].ln) //bbb

}

// anonymus struct
func main1() {

	//1
	e := struct {
		fn  string
		age int
	}{
		"aa", 30,
	}

	fmt.Println(e) //{aa 30}

	//2
	e1 := []struct {
		fn  string
		age int
	}{
		{"aa", 30},
		{"bb", 50},
		{"cc", 70},
	}

	fmt.Println(e1)                            //[{aa 30} {bb 50} {cc 70}]
	fmt.Println(e1[0].fn, e1[1].age, e1[2].fn) //aa 50 cc

	for i := 0; i < len(e1); i++ {
		fmt.Println(e1[i].fn, e1[i].age) //aa 30
	}
}

// anonymus field
type Emp struct {
	string
	int
}

func main2() {

	e1 := Emp{"aa", 10}
	e2 := &e1

	fmt.Println(e1) //{aa 10}
	fmt.Println(e2) //&{aa 10}

	fmt.Println((*e2).string, (*e2).int) //aa 10
	fmt.Println((e2).string, (e2).int)   //aa 10
	fmt.Println(e2.string, e2.int)       //aa 10
}
