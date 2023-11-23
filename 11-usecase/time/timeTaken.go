package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	log.Print("Initialized Refresh")
	t := time.Now()
	log.Print("Time Taken : ", time.Since(t))
}

//https://go.dev/play/p/hWu8KuW27ge