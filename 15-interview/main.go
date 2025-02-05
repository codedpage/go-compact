package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIResponse struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Record `json:"data"`
}

type Record struct {
	ID        int   `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Details   struct {
		Diagnosis     string `json:"diagnosis"`
		BloodPressure struct {
			Systolic  int `json:"systolic"`
			Diastolic int `json:"diastolic"`
		} `json:"blood_pressure"`
	} `json:"details"`
}

func main() {
	url := "https://jsonmock.hackerrank.com/api/medical_records?page=1"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Parse JSON response
	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Output the data
	fmt.Printf("Page: %d\n", apiResponse.Page)
	fmt.Printf("Total Records: %d\n", apiResponse.Total)
	fmt.Printf("Medical Records:\n")
	for _, record := range apiResponse.Data {
		fmt.Printf("ID: %d, Diagnosis: %s, Systolic: %d, Diastolic: %d\n",
			record.ID, record.Details.Diagnosis,
			record.Details.BloodPressure.Systolic,
			record.Details.BloodPressure.Diastolic)
	}
}
