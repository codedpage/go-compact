package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius = %0.2f, %s", e.radius, e.err)
	//return e.err
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero\n", err.radius) //Radius -20.00 is less than zero
			fmt.Println(err.err)                                       //radius is negative
			fmt.Println(err.Error())                                   //radius = -20.00, radius is negative
		}
		fmt.Println(err) //radius = -20.00, radius is negative
		return
	}
	fmt.Printf("Area of rectangle %0.2f", area)
}

//https://go.dev/play/p/3csVhGX1FCr
