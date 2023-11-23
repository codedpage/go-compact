package main

import "fmt"

func main1() {
	fmt.Println("Hello, 世界")

	a := make([]int, 5)
	fmt.Println(a)              //[0 0 0 0 0]
	fmt.Println(len(a), cap(a)) //5 5

	b := make([]int, 5, 5)
	fmt.Println(b)              //[0 0 0 0 0]
	fmt.Println(len(b), cap(b)) //5 5

	c := make([]int, 0, 5)
	fmt.Println(c)              //[]
	fmt.Println(len(c), cap(c)) //0 5

	a = append(a, 1)
	fmt.Println(a)              //[0 0 0 0 0 1]
	fmt.Println(len(a), cap(a)) //6 10

	b = append(b, 1)
	fmt.Println(b)              //[0 0 0 0 0 1]
	fmt.Println(len(b), cap(b)) //6 10

	c = append(c, 1)
	fmt.Println(c)              //[1]
	fmt.Println(len(c), cap(c)) //1 5
}

func main2() {

	a := make([]int, 4)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0] 4 4

	a = append(a, 1)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0 1] 5 8

	a = append(a, 2, 3, 4)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0 1 2 3 4] 8 8

	a = append(a, 5)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0 1 2 3 4 5] 9 16

	a = append(a, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0 1 2 3 4 5 6 7 8 9 10 11 12] 16 16

	a = append(a, 13)
	fmt.Println(a, len(a), cap(a)) //[0 0 0 0 1] 5 8

}

func main() {
	var slice []int = nil
	//var slice = make([]int, 0, 0)
	fmt.Println(slice, len(slice), cap(slice)) //[] 0 0

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice)) //[1] 1 1

	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice)) //[1 2] 2 2

	slice = append(slice, 3)
	fmt.Println(slice, len(slice), cap(slice)) //[1 2 3] 3 4

	slice = append(slice, 4)
	fmt.Println(slice, len(slice), cap(slice)) //[1 2 3 4] 4 4

	slice = append(slice, 5)
	fmt.Println(slice, len(slice), cap(slice)) //[1 2 3 4 5] 5 8

}
