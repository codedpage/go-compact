// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	a := []int{20, 10, 40, 50}

	fmt.Println(a)

	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println(a)
}
