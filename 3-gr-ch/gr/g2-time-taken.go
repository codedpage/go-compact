package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello(s string) {
	fmt.Println(s)
}

func main() {
	t := time.Now()
	go hello("11")
	go hello("22")
	go hello("33")
	go hello("44")
	go hello("55")

	time.Sleep(1 * time.Millisecond)

	fmt.Println("NumCPU :", runtime.NumCPU(), " | NumGoroutine :", runtime.NumGoroutine()) //NumCPU : 8  | NumGoroutine : 6
	fmt.Print("Time Taken : ", time.Since(t))
}
