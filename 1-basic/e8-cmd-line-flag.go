package main

import (
	"flag"
	"fmt"
)

func main() {

	// -word india -num 50 -fork=true
	// --word india --num 50 --fork=true
	w := flag.String("word", "foo", "a string")
	n := flag.Int("num", 10, "an int")
	f := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("word:", *w)
	fmt.Println("num:", *n)
	fmt.Println("fork:", *f)

}

func main1() {
	// -str delhi -val 100 tail
	var s string
	flag.StringVar(&s, "str", "bar", "a string var")

	var i int64
	flag.Int64Var(&i, "val", 20, "a string var")
	flag.Parse()

	fmt.Println("svar:", s)
	fmt.Println("ivar:", i)
	fmt.Println("tail:", flag.Args())
}
