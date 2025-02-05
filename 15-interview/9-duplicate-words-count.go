//"Go is awesome. Go is fast." angel = glean

package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "Go is awesome. Go is fast."

	b := map[string]int{}

	d := strings.Split(a, " ")

	for _, e := range d {

		//fmt.Println(e, len(e))
		//Output:

		b[e]++
	}

	fmt.Println(b)

}
