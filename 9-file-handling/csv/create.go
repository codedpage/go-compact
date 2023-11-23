package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.Create("create.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	t := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	randNum := fmt.Sprintf("%d", rand.Intn(100))
	var data = [][]string{{"Line1", t}, {"Line2", randNum}}

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
