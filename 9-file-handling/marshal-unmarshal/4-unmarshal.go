package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {

	//by string
	json_string := `
    {
        "firstname": "Rahul",
        "lastname": "Singh",
        "city": "Mumbai"
    }`

	emp := Employee{} //new(Employee)
	json.Unmarshal([]byte(json_string), &emp)
	fmt.Println(emp)

	boa()
}

func boa() {

	jsonstr := `{
		"Firstname" : "Bob",
		"Lastname": "Smith",
		"age" : 30
	}`

	type Employee struct {
		Firstname string
		Lastname  string
		Age       int
	}
	var emp Employee
	err := json.Unmarshal([]byte(jsonstr), &emp)

	if err == nil {
		fmt.Println(emp)
	}
}
