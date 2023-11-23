package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	json_string := []byte(`{"a":10, "b":20}`)

	//json-unmarshal
	type val struct {
		Aaa int `json:"a"`
		Bbb int `json:"b"`
	}

	//by val
	var c val //c := new(val)
	if err := json.Unmarshal(json_string, &c); err != nil {
		fmt.Printf("JSON unmarshaling failed - by val: %s", err)
	}
	fmt.Println(c.Aaa, c.Bbb) // 10, 20
	fmt.Println(c)            //{10 20}

}
