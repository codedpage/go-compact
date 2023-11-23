package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "strings"
)

//func
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

//=-----------------

var xmlUpdateStatus = false

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.htm")

	} else {
		r.ParseMultipartForm(10 << 20) // 10 * 1mb = 10mb

		file, handler, err := r.FormFile("xmlFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		xmlUpdateStatus = true

		s := bufio.NewScanner(file)
		c := 1
		for s.Scan() {

			fmt.Println(c, "===============", s.Text())
			c = c + 1
		}

		if xmlUpdateStatus {
			fmt.Fprintf(w, "Successfully Uploaded File\n")
		}
	}
}
