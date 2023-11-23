package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	f, _ := os.Open("basic.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		for k, v := range record {

			fmt.Print(k, "-", v, " ")
		}
		fmt.Println()
	}

}
