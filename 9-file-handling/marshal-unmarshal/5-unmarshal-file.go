package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "http://localhost/api.php" //{"firstname": "rahul", "lastname": "singh", "city": "mumbai"}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	json_string, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	//debugger
	//json_string = []byte(`{"firstname": "rahul", "lastname": "singh", "city": "delhi"}`)

	//json-unmarshal
	type Employee struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		City      string `json:"city"`
	}

	e := new(Employee)
	if err := json.Unmarshal(json_string, e); err != nil {
		fmt.Printf("JSON unmarshaling failed - by ref: %s", err)
	}
	fmt.Println(e.FirstName, e.LastName, e.City) // rahul singh mumbai
	fmt.Println(e)                               //&{rahul singh mumbai}

}
