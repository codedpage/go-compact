// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func show(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	go show(&wg)
	wg.Add(1)
	wg.Wait()
}
