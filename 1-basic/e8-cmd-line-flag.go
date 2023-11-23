package main

import (
	"flag"
	"fmt"
)

func main() {

	// -word india -num 50 -fork=true
	// --word india --num 50 --fork=true
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 10, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)

}

func main1() {
	// -str delhi -val 100 tail
	var strVar string
	flag.StringVar(&strVar, "str", "bar", "a string var")

	var intVar int64
	flag.Int64Var(&intVar, "val", 20, "a string var")
	flag.Parse()

	fmt.Println("svar:", strVar)
	fmt.Println("ivar:", intVar)
	fmt.Println("tail:", flag.Args())
}
