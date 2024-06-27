// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
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
	operation(1, 1, 3, "÷")
}

func operation(n int, n1 int, q int, op string) {

	//calculate 1st digit
	c := float64(n)
	max, min := 0.0, 0.0

	max = math.Pow(10, c) - 1.0
	min = math.Pow(10, c-1.0)

	mx := int(max)
	mn := int(min)

	//calculate 2nd digit

	c1 := float64(n1)
	max1, min1 := 0.0, 0.0

	max1 = math.Pow(10, c1) - 1.0
	min1 = math.Pow(10, c1-1.0)

	mx1 := int(max1)
	mn1 := int(min1)

	//Generate Operation
	for i := 0; i < q; i++ {
		x := rand.Intn(mx-mn) + mn
		y := rand.Intn(mx1-mn1) + mn1

		if y > x {
			x, y = y, x
		}

		fmt.Printf("%d %s %d = ?\n", x, op, y)

	}

}
