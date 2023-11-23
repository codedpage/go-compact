package main

import (
	"fmt"
)

func show(a int) {
	fmt.Println("a =", a)
}
func main() {
	a := 5
	defer show(a)
	a = 10
	fmt.Println("Before deferred a =", a)

}

// Before deferred a = 10
// a = 5
