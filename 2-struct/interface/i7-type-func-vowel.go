package main

import "fmt"

type VowelFinder interface {
	FindVowels() []rune
}

type custString string

func (cs custString) FindVowels() []rune {
	var v []rune

	for _, r := range cs {
		if r == 'a' || r == 'e' {
			v = append(v, r)
		}
	}

	return v
}

func main() {
	a := custString("apple")
	var v VowelFinder
	v = a
	fmt.Println(v.FindVowels())
}

//https://go.dev/play/p/mGCrSS4Euen
