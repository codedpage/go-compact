package main

import (
	"fmt"
	. "fmt"
)

func main() {
	a := [3]string{"aa", "bb", "cc"}

	//array reference
	arrPrint(&a) //Wrong way :  It'll work, but don't use.

	//slice
	b := a[:]
	slicePrint(b) //Right way : Use slice instead of array.

	//interface
	c := [4]interface{}{1, "hello", 2.5, nil}
	infPrint(c)

	//multi dimension
	d := [3][2]string{{"aa", "bb"}, {"mm", "nn"}, {"xx", "yy"}}
	multiDimension(d)
}

/////////////////////////////////

func arrPrint(a *[3]string) {
	for k, v := range a {
		Printf("%d .... %s\n", k, v)
	}
}

func slicePrint(a []string) {
	for k, v := range a {
		Printf("%d ---- %s\n", k, v)
	}
}

func infPrint(b [4]interface{}) {
	for k, v := range b {

		switch v.(type) {
		case string:
			fmt.Println(k, "=====", v.(string))

		case int:
			fmt.Println(k, "=====", v.(int))

		case float64:
			fmt.Println(k, "=====", v.(float64))

		case bool:
			fmt.Println(k, "=====", v.(bool))

		case nil:
			fmt.Println(k, "=====", nil)

		}
	}
}

func multiDimension(a [3][2]string) {

	for _, v := range a {
		//fmt.Print(v)
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}
}
