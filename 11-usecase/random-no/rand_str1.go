package main

import (
	"crypto/rand"
	"fmt"
 )

func main() {
 
    str := "12345abcde"

	id := make([]byte, 5)
	// Fill our array with random numbers
	rand.Read(id)
	fmt.Println(id)
	for i, b := range id {
		id[i] = str[b%10]
		fmt.Println(id[i],b%10);
	}

	fmt.Println(id, string(id))
}


/*[252 188 177 72 205]
51 2
100 8
99 7
51 2
97 5
[51 100 99 51 97] 3dc3a*/