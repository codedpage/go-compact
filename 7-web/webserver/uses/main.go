package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Load index file")
	http.ServeFile(w, r, "index.htm")
}

func save(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Load save file")
	http.ServeFile(w, r, "save.htm")
}

func setupRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/save", save)
	http.ListenAndServe(":8082", nil)
}

func main() {

	setupRoutes()
}
