package main

import (
	"fmt"
	"net/http"
)

func main() {

	main1()
	main2()
	main3()
	main4()
}

//HandleFunc - 1
func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HandleFunc - 1")
	})
	http.ListenAndServe(":3001", nil)
}

//HandleFunc -2
func main2() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3002", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HandleFunc - 2")
	//http.ServeFile(w, r, "index.hmt")
}

//Handle - HandlerFunc
func main3() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle - HandlerFunc")
	}))
	http.ListenAndServe(":3003", nil)
}

//Handle - http.Handler
func main4() {
	http.Handle("/", Save())
	http.ListenAndServe(":3004", nil)
}

func Save() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handle - http.Handler")
	})
}
