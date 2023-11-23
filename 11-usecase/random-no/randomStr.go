package main

import (
	"crypto/rand"
	"fmt"
)

func main1() {

	str := "12345abcde"

	id := make([]byte, 5)
	// Fill our array with random numbers
	rand.Read(id)
	fmt.Println(id)
	for i, b := range id {
		id[i] = str[b%10]
		fmt.Println(id[i], b%10)
	}

	fmt.Println(id, string(id))
}

func main() {

	const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const idSourceLen = byte(len(idSource))

	length := 5
	id := make([]byte, length)
	rand.Read(id)

	//debugger
	fmt.Println(id, idSourceLen, length)

	for i, b := range id {

		id[i] = idSource[b%idSourceLen]
		//debugger
		fmt.Println(id[i], "--", b%idSourceLen, " = ", string(id[i]))
	}

	fmt.Println(id)
	fmt.Println(string(id))

}

//https://go.dev/play/p/8JiWpys6ROu
