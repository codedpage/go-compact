package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define structures to match the JSON format
type APIResponse struct {
	Page       int             `json:"page"`
	PerPage    int             `json:"per_page"`
	Total      int             `json:"total"`
	TotalPages int             `json:"total_pages"`
	Data       []MedicalRecord `json:"data"`
}

type MedicalRecord struct {
	ID        int64     `json:"id"`
	Timestamp int64     `json:"timestamp"`
	Diagnosis Diagnosis `json:"diagnosis"`
	Vitals    Vitals    `json:"vitals"`
	Doctor    Doctor    `json:"doctor"`
	UserID    int       `json:"userId"`
	UserName  string    `json:"userName"`
	UserDOB   string    `json:"userDob"`
	Meta      Meta      `json:"meta"`
}

type Diagnosis struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Severity int    `json:"severity"`
}

type Vitals struct {
	BloodPressureDiastole int     `json:"bloodPressureDiastole"`
	BloodPressureSystole  int     `json:"bloodPressureSystole"`
	Pulse                 int     `json:"pulse"`
	BreathingRate         int     `json:"breathingRate"`
	BodyTemperature       float64 `json:"bodyTemperature"`
}

type Doctor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Meta struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

const url = "https://jsonmock.hackerrank.com/api/medical_records?page=1"

func main() {
	// Make an HTTP GET request
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

	// Parse JSON response
	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the parsed data
	fmt.Printf("Page: %d, Total Records: %d\n", result.Page, result.Total)
	for _, record := range result.Data {
		fmt.Printf("\nRecord ID: %d\n", record.ID)
		fmt.Printf("User: %s (DOB: %s)\n", record.UserName, record.UserDOB)
		fmt.Printf("Diagnosis: %s (Severity: %d)\n", record.Diagnosis.Name, record.Diagnosis.Severity)
		fmt.Printf("Doctor: %s\n", record.Doctor.Name)
		fmt.Printf("Vitals - BP: %d/%d, Pulse: %d, Breathing Rate: %d, Temp: %.1f\n",
			record.Vitals.BloodPressureSystole, record.Vitals.BloodPressureDiastole,
			record.Vitals.Pulse, record.Vitals.BreathingRate, record.Vitals.BodyTemperature)
		fmt.Printf("Height: %d cm, Weight: %d kg\n", record.Meta.Height, record.Meta.Weight)
	}
}
