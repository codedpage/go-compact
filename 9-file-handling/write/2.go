package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//1st file
	d1 := []byte("aaa\nbbb\n")
	err := os.WriteFile("./tmp/dat1.txt", d1, 0644)
	check(err)

	//2nd file
	f, err := os.Create("./tmp/dat2.txt")
	check(err)
	defer f.Close()

	//some
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//xxx
	n3, err := f.WriteString("xxx\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	//yyy
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("yyy\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
