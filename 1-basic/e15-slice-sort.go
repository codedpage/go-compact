package main

import (
	. "fmt"
	"sort"
)

func main() {

	strArray := [...]string{"bb", "cc", "aa", "dd", "ee"}
	str := strArray[:]

	s1 := str[2:4]
	Println("len", len(s1)) // 4-2 = 2
	Println("cap", cap(s1)) // 5-2 = 3
	Println(s1)             //s[2], s[3] : [cc dd]

	sort.Strings(str)
	Println("Strings:", str)

	ints := []int{30, 10, 20}
	sort.Ints(ints)
	Println("Ints:   ", ints)

	isSorted := sort.IntsAreSorted(ints)
	Println("Sorted: ", isSorted)

}

//https://go.dev/play/p/_FBUUGPqQC6
