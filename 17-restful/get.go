package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonmock.hackerrank.com/api/medical_records?page=1"

// Struct to parse JSON response
type APIResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"` // Use specific struct if needed
}

func main() {
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse JSON
	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the entire response
	fmt.Printf("API Response: %+v\n", result)
}
