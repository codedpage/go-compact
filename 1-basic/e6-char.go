package main

import "fmt"

func main() {
	fmt.Println('A') // 65
	fmt.Printf("%c\n", 8377)
	for r := 'A'; r <= 'z'; r++ {
		//fmt.Println(r, string(r))                            // 65 A
		fmt.Printf("%T %v %c %b %o %x \n", r, r, r, r, r, r) //int32 65 A 1000001 101 41
	}

}
