package main

import (
	. "fmt"
	"unicode/utf8"
)

func main() {

	var a string = "aaa bbb"
	Println(a, ": Length =", len(a))

	word1 := "Señor"
	Println(utf8.RuneCountInString(word1)) //5
	Println(len(word1))                    //6

	word2 := "Pets"
	Println(utf8.RuneCountInString(word2)) //4
	Println(len(word2))                    //4

	reverse1()
	reverse2()
	reverse3()
}

// by for-loop
func reverse1() {
	b := "welcome"
	c := make([]rune, len(b))
	for i, j := len(b)-1, 0; i >= 0; i, j = i-1, j+1 {
		c[j] = rune(b[i])
	}
	Println(string(c))
}

// by for-range
func reverse2() {
	b := "welcome"
	c := make([]rune, len(b))
	for i, v := range b {
		c[len(b)-i-1] = v
	}
	Println(string(c))
}

// by string - concatenation
func reverse3() {
	b := "welcome"
	s := ""
	for _, v := range b {
		s = string(v) + s
	}
	Println(s)
}
