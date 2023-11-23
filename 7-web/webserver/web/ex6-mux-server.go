package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		randNum := rand.Intn(100)
		fmt.Fprintf(w, "Hello mux server %d", randNum)

	}))
	http.ListenAndServe(":3000", mux)
}
