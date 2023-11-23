package main

import (
	. "fmt"
)

func main() {

	var a string = "aa"
	Printf("%s %T \n", a, a) //aa string
	Printf("%T %c %x %o %b %v\n", a[0], a[0], a[0], a[0], a[0], a[0])
	//uint8 a 61 141 1100001 97

	type redString string
	var b redString = "bb"
	Printf("%s %T \n", b, b) //aa main.redString

	byteSlice := []byte("Welcome")
	Printf("%d %T \n", byteSlice, byteSlice)

	uint8Slice := []uint8("Welcome")
	Printf("%d %T \n", uint8Slice, uint8Slice)

	runeSlice := []rune("Welcome")
	Printf("%d %T \n", runeSlice, runeSlice)

	int32Slice := []int32("Welcome")
	Printf("%d %T \n", int32Slice, int32Slice)

}
