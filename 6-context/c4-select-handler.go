package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/*
How to run

open two console
console - 1: For sever  :  server is ready

conosle - 2: For client :  hello controller, signal breaker (Ctl+c)
curl http://localhost:8091/hello
ctl+c
Control goes here => case <-ctx.Done():
*/
func main() {

	http.HandleFunc("/hello", hello)
	log.Println("Server Ready...")
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started $$$$$$$$$$$$$$")
	defer fmt.Println("server: hello handler ended ##########")

	select {
	case <-time.After(5 * time.Second):
		randomStr := fmt.Sprintf("Hello Web : %d", rand.Intn(100))
		fmt.Fprintf(w, randomStr+"\n")
		fmt.Println("gracefully exit ...")

	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("forcefully exit @@@")
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
