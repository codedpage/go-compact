package main

import . "fmt"

func main0() {

	a := map[string]int{
		"xx": 10,
		"yy": 20,
		"zz": 30,
	}

	for k, v := range a {
		Println(k, v)
	}

	if v, ok := a["zz"]; ok {
		Println(v, ok) //30 true
	}

	delete(a, "yy")
	Println(a) //map[xx:10 zz:30]

}

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["aa"] = 10
	Println(a) //map[aa:10]

	var d map[string]int = map[string]int{}
	d["dd"] = 40
	Println(d) //map[dd:40]

	c := map[string]int{}
	c["cc"] = 30
	Println(c) //map[cc:30]

	//most used
	b := make(map[string]int)
	b["aa"] = 10
	b["bb"] = 20
	Println(b) //map[aa:10 bb:20]

	e := map[string]int{"aa": 10, "bb": 20}
	Println(e) //map[aa:10 bb:20]

}
