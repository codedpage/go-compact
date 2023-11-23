package main

import (
	"fmt"
)

// 'a' - any valid unicode character within single quote is a rune

func mutate(s string) string {

	h := []rune(s)
	h[0] = 'a'
	return string(h)
}

func main() {
	str := "hello"
	fmt.Println(mutate(str))
}
