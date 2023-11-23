package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Load a TXT file.
	f, _ := os.Open("line-by-line.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record, "==>", len(record))
		for value := range record {
			fmt.Println(value, "-", record[value])
		}
		fmt.Println("###############")
	}
}
