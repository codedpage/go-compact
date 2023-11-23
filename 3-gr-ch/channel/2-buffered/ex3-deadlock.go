package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "Ram"
	ch <- "Kumar"
	//ch <- "Singh"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
