package main

import (
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

func (o order) String() string {
	return fmt.Sprintf("%d, %d", o.ordId, o.customerId)
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func main() {

	i := 10
	fmt.Printf("%d %T\n", i, i)

	o := order{
		ordId:      11,
		customerId: 22,
	}

	//String() automatically called
	fmt.Println(o) // {11 22} or 11, 22

	fmt.Println(createQuery(o))
}
