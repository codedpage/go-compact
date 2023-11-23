package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Config struct {
		Name     string `json:"SiteName"`
		URL      string `json:"SiteUrl"`
		Database struct {
			Name     string
			Host     string
			Port     int
			Username string
			Password string
		}
	}
	conf := Config{}
	data, err := os.ReadFile("./unmarshaling-config.json")
	if err != nil {
		panic(err)
	}

	jsonString := `{ "SiteName": "My Blog", "SiteUrl": "www.abc.com", "Database": { "Name": "cats", "Host": "localhost", "Port": 3306, "Username": "admin", "Password": "admin" } } `
	data = []byte(jsonString)

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	//convert in interface : readable data
	fmt.Println(conf)
	fmt.Println("=========")

	//formatting
	fmt.Printf("%s | %s | %s\n", conf.Name, conf.URL, conf.Database.Host)
	db := conf.Database
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s\n",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

}
