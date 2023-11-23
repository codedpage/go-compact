package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//func
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3001", nil)
}

//=-----------------

var xmlUpdateStatus = false

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.htm")

	} else {
		r.ParseMultipartForm(10 << 20)

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

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("tmp", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!

		if xmlUpdateStatus {
			fmt.Fprintf(w, "Successfully Uploaded File\n")
		}
	}
}
