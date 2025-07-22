package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {

	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)

	fmt.Println(string(byteValue))

	for _, v := range users.Users {
		fmt.Println(v.Name, v.Type, v.Social.Facebook, v.Social.Twitter, v.Social.Youtube)
	}

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("Type: " + users.Users[i].Type)
		fmt.Println("Name: " + users.Users[i].Name)
		fmt.Println("Social: " + users.Users[i].Social.Facebook)
	}
}
