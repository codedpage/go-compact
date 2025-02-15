package main

import "fmt"

// Defining an empty struct
type Logger struct{}

// Defining a method on the empty struct
func (l Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func main() {
	l := Logger{} // No memory overhead
	l.Log("This is a log message")
}
