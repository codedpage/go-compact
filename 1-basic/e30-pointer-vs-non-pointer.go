package main

import "fmt"

// ---------- 1. int (Non-pointer wrap) ----------
func updateInt(x int) {
	x = 10
}

// ---------- 2. array (Non-pointer wrap) ----------
func updateArray(a [3]int) {
	a[0] = 100
}

// ---------- 3. struct (Non-pointer wrap) ----------
type User struct {
	Name string
}

func updateStruct(u User) {
	u.Name = "Manoj"
}

// ---------- 4. slice (Mixed wrap) ----------
func updateSlice(s []int) {
	s[0] = 200
}

// ---------- 5. map (Pointer-like wrap) ----------
func updateMap(m map[string]int) {
	m["age"] = 40
}

// ---------- 6. channel (Pointer-like wrap) ----------
func sendValue(ch chan int) {
	ch <- 50
}

// ---------- 7. struct pointer (Pointer wrap) ----------
type Person struct {
	Name string
}

func updateStructPtr(p *Person) {
	p.Name = "Manoj"
}

func main() {

	// 1. int
	a := 5
	updateInt(a)
	fmt.Println("int:", a) // 5 ❌

	// 2. array
	arr := [3]int{1, 2, 3}
	updateArray(arr)
	fmt.Println("array:", arr) // [1 2 3] ❌

	// 3. struct
	u := User{Name: "Ram"}
	updateStruct(u)
	fmt.Println("struct:", u.Name) // Ram ❌

	// 4. slice
	s := []int{10, 20, 30}
	updateSlice(s)
	fmt.Println("slice:", s) // [200 20 30] ✅

	// 5. map
	m := map[string]int{"age": 30}
	updateMap(m)
	fmt.Println("map:", m) // map[age:40] ✅

	// 6. channel
	ch := make(chan int)
	go sendValue(ch)
	fmt.Println("channel:", <-ch) // 50 ✅

	// 7. struct pointer
	p := Person{Name: "Ram"}
	updateStructPtr(&p)
	fmt.Println("struct pointer:", p.Name) // Manoj ✅
}

// int: 5
// array: [1 2 3]
// struct: Ram
// slice: [200 20 30]
// map: map[age:40]
// channel: 50
// struct pointer: Manoj
