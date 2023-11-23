package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	//create
	file, err := os.OpenFile("dw.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	fmt.Println(err, file)
	//panic("exit")

	//apend
	f, err := os.OpenFile("dw.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprint(f, time.Now(), "\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
