package main

import (
	. "fmt"
)

func main() {

	//slice + value
	s := []string{"aa", "bb", "cc"}
	s = append(s, "dd")
	Println(s)

	//slice + slice
	a := []string{"aa", "bb"}
	b := []string{"xx", "yy"}

	c := append(a, b...)
	Println(c)

	//slice + values
	d := append([]string{}, "xx", "yy", "zz")
	e := append(d, s...)
	Println(e)
}
