package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	input :=
		`aa bb cc dd ee ff
mm nn
xx yy`

	f := strings.NewReader(input)
	scanner := bufio.NewScanner(f)
	//	scanner.Split(bufio.ScanWords)

	fmt.Println(scanner)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
