package main

import "fmt"

func main() {

	m := map[int]string{} //m := make(map[int]string)
	m[1] = "aa"
	m[2] = "bb"
	m[3] = "cc"
	getMap(m)

	m2 := map[int]string{
		1: "aaa",
		2: "bbb",
		3: "ccc",
	}
	getMap(m2)

	m3 := map[int]map[int]string{
		1: map[int]string{10: "a", 20: "b", 30: "c"},
		2: map[int]string{10: "a", 20: "b", 30: "c"},
		3: map[int]string{10: "a", 20: "b", 30: "c"},
	}
	fmt.Println(m3)

	/*
		country -> state -> city ->
		c1 -> s(1,2,3) -> c(11, 12, 13,) (21, 22, 23) (31, 32, 33)

	*/

	m4 := map[int]map[int]map[int]string{
		1: map[int]map[int]string{10: map[int]string{10: "a", 20: "b", 30: "c"}},
	}
	fmt.Println(m4)

	m5 := map[int]map[int]map[int]map[int]string{1: map[int]map[int]map[int]string{1: map[int]map[int]string{10: map[int]string{10: "a", 20: "b", 30: "c"}}}}
	fmt.Println(m5)

}

func getMap(m map[int]string) {

	ki := []interface{}{}
	vs := []interface{}{}

	for k, v := range m {
		ki = append(ki, k)
		vs = append(vs, v)
	}

	fmt.Println(ki)
	fmt.Println(vs)

}

//https://go.dev/play/p/ZXvlQ9AWyQ4