package main

import (
	"fmt"
	"sort"
)

func main() {

	//s := []int{5, 2, 6, 3, 1, 4}
	//s1 := sort.IntSlice(s)

	//s := []string{"dd", "aa", "cc", "ee", "bb"}
	//s1 := sort.StringSlice(s)

	s := []float64{5.1, 2.2, 6.3, 3.4, 1.5, 4.6}
	s1 := sort.Float64Slice(s)

	//copy
	original := make([]float64, len(s))
	copy(original, s)

	//Reverse
	s2 := sort.Reverse(s1)

	//s2 := s1
	sort.Sort(s2)

	fmt.Println(original)
	fmt.Println(s)

}
//https://go.dev/play/p/lvWSsCenUIg
