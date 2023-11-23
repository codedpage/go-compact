package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

/*
How to run

open two console
console - 1: For sever  :  server is ready
conosle - 2: For client :
curl http://localhost:8000
ctl+c
Control goes here => case <-ctx.Done():
*/

func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Fprint(os.Stdout, "Processing request...\n") //console

		select {
		case <-time.After(15 * time.Second):
			w.Write([]byte("Request Processed\n")) //webpage

		case <-ctx.Done(): //cancelled
			fmt.Fprint(os.Stderr, "Request cancelled\n") //console
		}
	}))
}
