package main

import (
	"encoding/json"
	"fmt"
)

type Dc map[string]interface{}
type Extra map[string]interface{}

type Metadata struct {
	NetworkType string `json:"network_type"`
	Dc          Dc
	Fabric      string
}

type FabricInfo struct {
	Id         string
	Name       string
	Metadata   Metadata
	Subscribed bool
}

func main() {
	fmt.Println("Hello, 世界")

	data := []byte(`{
		"id": "123456",
		"name": "test data",
		"metadata": {
			"network_type": "test",
			"dc": {
				"test": "test",
				"isPhone": true,
				"extra": {
					"email": "a@b.com"
				}
			},
			"fabric": "test"
		},
		"subscribed": false
	}`)

	var fab FabricInfo
	err := json.Unmarshal(data, &fab)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(fab.Metadata.NetworkType)
	fmt.Println(fab.Id)
	email, ok := fab.Metadata.Dc["extra"].(map[string]interface{})
	if ok {
		fmt.Println(email["email"])
	} else {
		fmt.Println("Email not found")
	}
}
