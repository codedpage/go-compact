package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func orderObject(q interface{}) {
	v := reflect.ValueOf(q)
	t := reflect.TypeOf(q)

	k := reflect.TypeOf(q).Kind()
	k1 := reflect.ValueOf(q).Kind()

	fmt.Println("Value ", v)
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	fmt.Println("Kind ", k1)
	fmt.Println(q)

	fmt.Println("Name ", reflect.TypeOf(q).Name())

}
func main() {
	o := order{
		ordId:      1001,
		customerId: 201,
	}
	orderObject(o)

}

// Value  {1001 201}
// Type  main.order
// Kind  struct
// Kind  struct
// {1001 201}
// Name  order
