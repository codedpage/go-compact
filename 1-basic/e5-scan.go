package main

import "fmt"

func main() {

	var n int
	fmt.Printf("Enter the no: ")

	//fmt.Scanf("%d", &n)
	//fmt.Scan(&n)
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}

// pointer
func main1() {

	n := new(int)
	fmt.Printf("Enter the no: ")

	fmt.Scanln(n)

	if *n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
