package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/dev"))))

	fmt.Println(http.ListenAndServe(":3002", mux))
}

func main1() {
	port := flag.String("p", "3000", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	handler := http.FileServer(http.Dir(*directory))
	http.Handle("/", handler)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
