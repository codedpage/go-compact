package main

import "fmt"

type rectangle struct {
	l, w int
}

func (r rectangle) area()   { fmt.Println(r.l * r.w) }
func (r *rectangle) areap() { fmt.Println(r.l * r.w) }

func main() {
	r := rectangle{
		l: 10,
		w: 5,
	}

	r.area()    //v-v
	(&r).area() //r-v

	r.areap()    //v-r
	(&r).areap() //r-r

	//pointer variable
	p := &r
	p.area()  //r-v
	p.areap() //r-r

}

//https://go.dev/play/p/eLfXpmnvnBd
