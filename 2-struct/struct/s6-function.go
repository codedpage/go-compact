package main

import "fmt"

type rectangle struct {
	l, w int
}

func area(r rectangle)   { fmt.Println(r.l * r.w) }
func areap(r *rectangle) { fmt.Println(r.l * r.w) }

func main() {
	r := rectangle{
		l: 10,
		w: 5,
	}

	//v-v
	area(r)

	//r-r
	areap(&r)

	//pointer variable
	p := &r
	area(*p) //v-v
	areap(p) //r-r

}

//https://go.dev/play/p/cWBzCPt_l-1
