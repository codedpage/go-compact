package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	randNo := rand.Intn(max-min+1) + min
	fmt.Println(randNo)
}

//https://go.dev/play/p/TYHfzjJzVPu
