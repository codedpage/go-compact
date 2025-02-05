//angel = glean

package main

import (
	"fmt"
)

func main() {

	a := "glean"
	b := "angeli"

	fmt.Println(a)
	fmt.Println(b)

	m := map[string]string{}

	for _, e := range a {
		fmt.Println(string(e))
		m[string(e)] = string(e)
	}

	fmt.Println(m)

	var flag bool
	flag = true
	for _, e := range b {

		// if v, f := m[string(e)]; f {
		// 	fmt.Println(v, f, string(e))
		// } else {
		// 	flag = false
		// 	fmt.Println(v, f, string(e))
		// }

		if v, f := m[string(e)]; !f {
			flag = false
			fmt.Println(v, f, string(e))
		}
	}

	if flag {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not Anagram")
	}

}
