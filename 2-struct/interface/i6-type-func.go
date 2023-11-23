package main

import "fmt"

type customFloat64 float64

func (m customFloat64) Test() {
	fmt.Println(m)
}

type Tester interface {
	Test()
}

func Describe(t Tester) {
	fmt.Printf("%T ------ %v \n", t, t)
}

func main() {
	var t Tester
	f := customFloat64(10.5)
	t = f

	t.Test()
	Describe(t)

}

//https://go.dev/play/p/vKJ4QFtOmUQ
