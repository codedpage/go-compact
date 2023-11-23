// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	c := 1
	for country := range countryCapitalMap {
		ele := []byte(country)
		fn := fmt.Sprintf("file_%d%d.txt", c, rand.Intn(5))
		println(fn)
		os.WriteFile(fn, ele, 0644)
		c++
	}

}
