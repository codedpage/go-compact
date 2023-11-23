package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("main.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	csvWriter := csv.NewWriter(file)

	z := append([]string{"Line1", "aa", "bb"}, []string{"cc"}...)

	//Single
	csvWriter.Write(z)
	fmt.Println(z)

	//Multiple
	strWrite := [][]string{{"Line2", "dd", "ee", "ff"}, z, {"Line3", "xx", "yy", "zz"}}
	csvWriter.WriteAll(strWrite)
	fmt.Println(strWrite)

	csvWriter.Flush()
}
