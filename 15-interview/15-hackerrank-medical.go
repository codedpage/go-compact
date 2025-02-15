package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIResponse struct {
	Page       int             `json:"page"`
	PerPage    int             `json:"per_page"`
	Total      int             `json:"total"`
	TotalPages int             `json:"total_pages"`
	Data       []PatientRecord `json:"data"`
}

type Diagnosis struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Severity int    `json:"severity"`
}

type Vitals struct {
	BloodPressureDiastole int         `json:"bloodPressureDiastole"`
	BloodPressureSystole  int         `json:"bloodPressureSystole"`
	Pulse                 int         `json:"pulse"`
	BreathingRate         int         `json:"breathingRate"`
	BodyTemperature       interface{} `json:"bodyTemperature"`
}

type Doctor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Meta struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type PatientRecord struct {
	ID        int       `json:"id"`
	Timestamp int64     `json:"timestamp"`
	Diagnosis Diagnosis `json:"diagnosis"`
	Vitals    Vitals    `json:"vitals"`
	Doctor    Doctor    `json:"doctor"`
	UserID    int       `json:"userId"`
	UserName  string    `json:"userName"`
	UserDob   string    `json:"userDob"`
	Meta      Meta      `json:"meta"`
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
			record.ID, record.Diagnosis.Name,
			record.Vitals.BloodPressureSystole,
			record.Vitals.BloodPressureDiastole)
		fmt.Println(record, "\n")

	}

}
