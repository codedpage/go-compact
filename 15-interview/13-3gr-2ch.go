// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	c2 := make(chan string)

	go f1("abc", c)

	go f2(c, c2)
	go f3(c2)

	time.Sleep(1 * time.Second)
}

func f1(s string, c chan string) {

	c <- s
}

func f2(c chan string, c2 chan string) {
	c2 <- <-c
}

func f3(c2 chan string) {
	fmt.Println(<-c2)
}
