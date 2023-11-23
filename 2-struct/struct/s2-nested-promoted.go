package main

import "fmt"

type Address struct{ city, state string }
type Contact struct{ email, mobile string }
type Flat struct{ city string }

type Person struct {
	name    string
	age     int
	Address         //promoted
	contact Contact //nested
	flats   []Flat  //nested
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20

	//promoted
	p.Address = Address{
		city:  "ND",
		state: "Delhi",
	}

	//nested
	p.contact.email = "user@gmail.com"
	p.contact.mobile = "9000000000"
	p.flats = []Flat{
		{"Delhi"}, {city: "Mumbai"},
	}

	fmt.Println(p)               //{aa 20 {ND Delhi} {user@gmail.com 9000000000} [{Delhi} {Mumbai}]}
	fmt.Println(p.contact.email) //user@gmail.com

	/////////////////////
	p1 := Person{
		name: "bb",
		age:  30,
		Address: Address{ //promoted
			city:  "New Delhi",
			state: "Delhi",
		},

		contact: Contact{ //nested
			"user2@gmail.com", "9000000002",
		},

		flats: []Flat{
			{"Delhi"}, {"Mumbai"}, {"Goa"},
		},
	}

	fmt.Println(p1.name, p1.age)                     //bb 30
	fmt.Println(p1.city, p1.state)                   //promoted 	=> New Delhi Delhi
	fmt.Println(p1.contact.email, p1.contact.mobile) //nested 		=> user2@gmail.com 9000000002
	fmt.Println(p1.flats, p1.flats[0])               //[{Delhi} {Mumbai} {Goa}] {Delhi}
}
