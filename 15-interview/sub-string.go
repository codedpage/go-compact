// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	s := "acsabcdefdkl"
	fmt.Println(s)

	a := map[string]int{}

	r := ""

	for i := 0; i < len(s); i++ {

		if _, k := a[string(s[i])]; k {

			a[string(s[i])]++

			continue
		}
		r = r + string(s[i])
	}

	fmt.Println(r)

}
