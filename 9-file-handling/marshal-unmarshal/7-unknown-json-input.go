package main

import (
	"encoding/json"
	"fmt"
)

func FooJSON(input string) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}
	val, ok := data["val"]
	if !ok {
		fmt.Println("Key 'val' not found in the input.")
		return
	}
	switch v := val.(type) {
	case float64:
		if v == float64(int(v)) {
			fmt.Printf("Int %d\n", int(v))
		} else {
			fmt.Printf("Float %f\n", v)
		}
	case string:
		fmt.Printf("String %s\n", v)
	default:
		fmt.Printf("Something else\n")
	}
}

func main() {
	FooJSON(`{"val": 123.5}`)
	FooJSON(`{"val": 20}`)
	FooJSON(`{"val": "bar"}`)
	FooJSON(`{"val": []}`)
}
