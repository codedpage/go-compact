package main

import "fmt"

func main() {

	var rs = map[string]string{
		"0": "aa",
		"1": "bb",
		"2": "cc",
	}
	fmt.Println(rs["1"])

	/////////////////////////////////
	var ds = map[int]string{
		0: "aa",
		1: "bb",
		2: "cc",
	}
	fmt.Println(ds[2])

	///////////////////////
	var es = map[int]string{0: "aa", 1: "bb", 2: "cc"}
	fmt.Println(es[0])

	strDict := map[string]string{"aa": "aa", "bb": "bb", "cc": "cc"}

	if value, ok := strDict["dd"]; !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Key found value is: ", value)
	}

}

//https://go.dev/play/p/ZEQnjPVnvmC
