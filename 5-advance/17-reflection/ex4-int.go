package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	aa := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", aa, aa)

	b := "aa"
	bb := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", bb, bb)

	///////////
	var x int = 5
	var x64 int64 = int64(x)
	fmt.Println(reflect.TypeOf(x))   //int
	fmt.Println(reflect.TypeOf(x64)) //int64
	fmt.Println(x, x64)              //5 5

	///////////
	var y interface{} = int64(x)
	var y64 int64 = y.(int64)
	fmt.Println(reflect.TypeOf(y))   //int64
	fmt.Println(reflect.TypeOf(y64)) //int64
	fmt.Println(y, y64)              //5 5

}
