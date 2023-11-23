package main

import (
	. "fmt"
)

/*
	lo = n
	hi = m
	len = m - n
	cap = len(a) - n
	values = a[n],a[n+1]... a[m-1]

case -1 : m > n , len = m-n, cap = len(a) - n
case -2 : m = n, len = 0, cap = len(a) - n
case -3 : m = n + 1, len = 1, cap = len(a) - n
*/
func main() {
	Println("Hello, playground")

	a := []int{1, 2, 3, 4, 5}

	//n:n blank values =>  []
	Println("m=n ", a[0:0], a[1:1], a[2:2], a[3:3], a[4:4], a[5:5])

	//n:n+1, single value =>  a[n]
	Println("m=n+1 ", a[0:1], a[1:2], a[2:3], a[3:4], a[4:5])

	//n:n+x, values => n to n+x-1
	Println("2:5 ", a[2:5]) // 3 4 5
	Println("1:4 ", a[1:4]) // 2 3 4
	Println("2:3 ", a[2:3]) // 3

	Println("1st Element ", a[:1], a[0:1])
	Println("Last Element ", a[len(a)-1:], a[len(a)-1:len(a)])

	//len : m-n
	//cap : len(array) - n
	n, m := 1, 4
	b := a[n:m]
	Println("Length :", len(b), b, m-n)     //3
	Println("Capacity :", cap(b), len(a)-n) //4
}
