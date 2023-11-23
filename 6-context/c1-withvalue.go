package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
)

// ex-1
func Operation1(ctx context.Context, n int) {
	log.Println("Operation1", ctx.Value(keyID), n)
	Operation2(ctx, n)
}

func Operation2(ctx context.Context, n int) {
	log.Println("Operation2", ctx.Value(keyID), n)
}

// we need to set a key that tells us where the data is stored
const keyID = "id"

func main() {

	//ex-1
	randNum := rand.Intn(100)
	ctx := context.WithValue(context.Background(), keyID, randNum)
	Operation1(ctx, randNum)

	//ex-2
	ctx2 := context.Background()
	ctx2 = middleLayer(ctx2)
	doSomethingCool(ctx2)
}

// ex-2
func middleLayer(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", 1001)
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
}
