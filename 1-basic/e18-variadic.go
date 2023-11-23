package main

import (
	. "fmt"
)

func main() {

	s := []string{"aa", "bb"}
	show(s...)

}

func show(s ...string) {
	Println(s)
}
