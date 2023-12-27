package main

import (
	"fmt"
	"sync"
)

func calculateSubarraySum(subarray []int, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	str := ""
	for _, num := range subarray {
		str = str + fmt.Sprintf("%d ", num)
		sum += num
	}

	resultChan <- fmt.Sprintf("%d = %s", sum, str)
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	subarraySize := len(arr) / 5

	resultChan := make(chan string, 5)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		startIndex := i * subarraySize
		endIndex := (i + 1) * subarraySize

		subarray := arr[startIndex:endIndex]

		go calculateSubarraySum(subarray, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	for partialSum := range resultChan {
		fmt.Printf("Sum: %s\n", partialSum)
	}
}
