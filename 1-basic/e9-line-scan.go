package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// once
func main0() {

	fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	//fmt.Println(input)
	fmt.Println("Input string: ", input.Text())
}

// uses -1
func main1() {
	fmt.Print("Enter text: ")
	c := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Input string: ", input.Text())
		fmt.Print("Enter text: ")
		c[input.Text()]++

		if input.Text() == "break" {
			break
		}
	}

	fmt.Println(c)
}

// uses -2
func main() {
	fmt.Printf("Enter the no: ")
	input := bufio.NewScanner(os.Stdin)
	c := make(chan int)
	for input.Scan() {
		i, _ := strconv.Atoi(input.Text())
		go sum(i, c)
		fmt.Println(<-c)
	}
}

func sum(i int, c chan int) {
	c <- i + 10
}
