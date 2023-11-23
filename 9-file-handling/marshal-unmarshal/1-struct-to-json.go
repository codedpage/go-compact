package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Employee struct {
		Name    string `json:"first_name"`
		Address string `json:"address,omitempty"`
		Draft   bool
	}

	//case-1
	e := Employee{"Ram", "Delhi", false}
	data, err := json.MarshalIndent(e, "", "  ") //formatted
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//case-2
	f := Employee{"Krishna", "", true}
	data1, err := json.Marshal(f) //unformatted string
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data1))

}
