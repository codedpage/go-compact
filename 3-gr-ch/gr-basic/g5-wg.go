// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func show(i int32, wg *sync.WaitGroup) {

	fmt.Println(i)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	go show(1, &wg)
	wg.Add(1)

	go show(2, &wg)
	wg.Add(1)

	go show(3, &wg)
	wg.Add(1)

	go show(4, &wg)
	wg.Add(1)

	wg.Wait()
}
