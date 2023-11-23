package main

import (
	"fmt"
)

type Author struct {
	firstName string
	lastName  string
	bio       string
}

func (a Author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type Post struct {
	title   string
	content string
	Author  //promoted
}

func (p Post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

//----------------------

func main() {
	author := Author{
		"Manoj",
		"Kumar",
		"Golang SSE",
	}
	post := Post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author,
	}
	post.details()
}

//https://go.dev/play/p/2aRsK4JjYp7
