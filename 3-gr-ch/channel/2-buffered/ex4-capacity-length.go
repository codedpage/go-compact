//The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make function.
//The length of the buffered channel is the number of elements currently queued in it.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "Rahul"
	ch <- "kumar"
	fmt.Printf("capacity = %d length = %d \n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("capacity = %d length = %d \n", cap(ch), len(ch))
}
