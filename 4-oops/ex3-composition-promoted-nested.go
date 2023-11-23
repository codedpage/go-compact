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

type Website struct {
	posts []Post //nested - slice
}

type Webpage struct {
	post Post //nested - single
	//Post      //promoted
}

func (w Website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

//-----------------------------------------------------

func main() {
	author1 := Author{
		"Manoj",
		"kumar",
		"SSE",
	}
	post1 := Post{
		"aa",
		"aaa",
		author1,
	}
	post2 := Post{
		"bb",
		"bbb",
		author1,
	}
	post3 := Post{
		"cc",
		"ccc",
		author1,
	}
	w := Website{
		posts: []Post{post1, post2, post3},
	}
	w.contents()

	wp := Webpage{
		post: post1,
		//post1, //promoted
	}
	fmt.Println(wp)
	fmt.Println(wp.post, wp.post.Author)
}
