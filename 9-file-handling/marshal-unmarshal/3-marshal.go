package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {

	enc := json.NewEncoder(os.Stdout)

	//by value
	emp1 := Employee{"Ram", "kumar", "Delhi"}

	//struct-to-json
	fmt.Println(emp1)                //{Ram kumar Delhi}
	enc.Encode(emp1)                 //{"firstname":"Ram","lastname":"kumar","city":"Delhi"}
	jsonStr, _ := json.Marshal(emp1) //using Marshal
	fmt.Println(jsonStr)             // //[123 34 102 105 114 ... 105 34 125]
	fmt.Println(string(jsonStr))     //{"firstname":"Ram","lastname":"kumar","city":"Delhi"}

	emp2 := Employee{}
	json.Unmarshal(jsonStr, &emp2)
	fmt.Println(emp2) //{Ram kumar Delhi}

	emp3 := new(Employee)
	json.Unmarshal(jsonStr, emp3)
	fmt.Println(emp3) //&{Ram kumar Delhi}

	//map-to-json
	d := map[string]string{"firstname": "Ram", "lastname": "kumar", "city": "Delhi"}
	enc.Encode(d)

	e := map[string]int{"a": 5, "b": 10}
	enc.Encode(e)

}
