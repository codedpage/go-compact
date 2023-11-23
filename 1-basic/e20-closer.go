package main

import . "fmt"

//Closures : a anonymous functions which can access the variables defined outside of function

func main() {

	n := 5
	func() {
		Println("n =", n)
	}()

	a := intSeq()
	Println(a(), a(), a()) //1 2 3

	b := intSeq()
	Println(b(), b()) //1 2

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
