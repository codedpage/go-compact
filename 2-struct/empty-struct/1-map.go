/*
Understanding Empty Struct (struct{}) in Golang with Examples
An empty struct (struct{}) is a struct type that has no fields.
It is useful when you need to define a type but don’t need to store any data.

🔹 Key Features of struct{}
Zero Memory Overhead: It occupies 0 bytes in memory.
Lightweight and Efficient: Since it has no fields, it avoids unnecessary memory allocation.
Common Use Cases:
Used for signaling (e.g., in channels).
Used for set-like behavior in maps (where only keys matter).
Used for implementing methods without data (like dummy types).
*/

package main

import "fmt"

func main() {
	// Define a set using map with struct{}
	set := make(map[string]struct{})

	// Adding values to the set
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["cherry"] = struct{}{}

	// Checking if a value exists in the set
	if _, exists := set["banana"]; exists {
		fmt.Println("Banana is in the set!")
	}

	// Iterating over the set
	fmt.Println("Fruits in the set:")
	for key := range set {
		fmt.Println(key)
	}
	fmt.Println(set) //map[apple:{} banana:{} cherry:{}]
}
