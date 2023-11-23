package main

import (
	"fmt"
	"net/http"
)

// ex1 : basic
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gopher...")
	})

	http.ListenAndServe(":3000", nil)
}

// ex2 : using handle
func main1() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello Gopher through function...")

	}))
	http.ListenAndServe(":3000", nil)
}

// ex3
func main2() {
	http.Handle("/", MyHandleFunction())
	http.ListenAndServe(":3000", nil)
}

func MyHandleFunction() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello Gopher through function...")

	})
}
