// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("\nAddition")
	operation(1, 1, 5, "+")

	fmt.Println("\nSubtraction")
	operation(1, 1, 5, "-")

	fmt.Println("\nMultiplication")
	operation(1, 1, 5, "X")

	fmt.Println("\nDivision")
	operation2(1, 1, 5, "÷")
}

func operation2(n int, n1 int, q int, op string) {

	//Generate Operation
	for i := 0; i < q; i++ {
		x := rand.Intn(30)
		y := rand.Intn(2) + 2

		fmt.Printf("%d %s %d = ?\n", x, op, y)

	}

}

func operation(n int, n1 int, q int, op string) {

	//Generate Operation
	for i := 0; i < q; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		if y > x {
			x, y = y, x
		}

		fmt.Printf("%d %s %d = ?\n", x, op, y)

	}

}
