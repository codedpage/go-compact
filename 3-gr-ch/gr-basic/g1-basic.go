package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {

	go hello()
	fmt.Println("main")
	time.Sleep(1*time.Second)
}

//https://go.dev/play/p/3d3xf9TqpsN
