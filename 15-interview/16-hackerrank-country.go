package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Country `json:"data"`
}

type Country1 struct {
	Name       string      `json:"name"`
	Population interface{} `json:"population"`
}

type Country2 struct {
	Name           string            `json:"name"`
	NativeName     string            `json:"nativeName"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2Code"`
	NumericCode    string            `json:"numericCode"`
	Alpha3Code     string            `json:"alpha3Code"`
	Currencies     []string          `json:"currencies"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Relevance      string            `json:"relevance"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Language       []string          `json:"language"`
	Languages      []string          `json:"languages"`
	Translations   map[string]string `json:"translations"`
	Population     int               `json:"population"`
	Latlng         []float64         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Borders        []string          `json:"borders"`
	Area           float64           `json:"area"`
	Gini           float64           `json:"gini"`
	Timezones      []string          `json:"timezones"`
}

type Country struct {
	Name           interface{} `json:"name"`
	NativeName     interface{} `json:"nativeName"`
	TopLevelDomain interface{} `json:"topLevelDomain"`
	Alpha2Code     interface{} `json:"alpha2Code"`
	NumericCode    interface{} `json:"numericCode"`
	Alpha3Code     interface{} `json:"alpha3Code"`
	Currencies     interface{} `json:"currencies"`
	CallingCodes   interface{} `json:"callingCodes"`
	Capital        interface{} `json:"capital"`
	AltSpellings   interface{} `json:"altSpellings"`
	Relevance      interface{} `json:"relevance"`
	Region         interface{} `json:"region"`
	Subregion      interface{} `json:"subregion"`
	Language       interface{} `json:"language"`
	Languages      interface{} `json:"languages"`
	Translations   interface{} `json:"translations"`
	Population     interface{} `json:"population"`
	Latlng         interface{} `json:"latlng"`
	Demonym        interface{} `json:"demonym"`
	Borders        interface{} `json:"borders"`
	Area           interface{} `json:"area"`
	Gini           interface{} `json:"gini"`
	Timezones      interface{} `json:"timezones"`
}

func fetchCountries(search string, population int) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries/?s=%s&p=%d", search, population)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	jsonData, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonData))
}

func main() {
	s := "us"   // Example dynamic search term
	p := 100090 // Example dynamic population filter

	fetchCountries(s, p)
}
