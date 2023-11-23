package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Load index file")
	http.ServeFile(w, r, "index.htm")
}

func save() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "save.htm")
	})
}

func mainOne() {

	//HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.htm")
	})

	//Handle-HandlerFunc
	http.Handle("/save", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "save.htm")
	}))

	http.ListenAndServe(":8080", nil)
}

func main() {

	//HandleFunc
	http.HandleFunc("/", index)
	//http.Handle("/", http.HandlerFunc(index))	

	//Handle-Handler-HandlerFunc
	http.Handle("/save", save())
	

	http.ListenAndServe(":8083", nil)
}
