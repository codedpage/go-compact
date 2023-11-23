package main

import (
	"context"
	"fmt"
	"time"
)

// main work
func operation2(ctx context.Context, i int) {

	select {
	// case <-time.After(1500 * time.Millisecond):
	// 	fmt.Println(ctx, i)
	case <-ctx.Done():
		fmt.Println("operation2 halted because operation1 returns error")
	default:
		fmt.Println(ctx, i)
	}
}

// If this operation returns an error cancel all operations using this context
func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	// Run operation2 with the same context we use for operation1

	for i := 1; i <= 10; i++ {
		go operation2(ctx, i)
		time.Sleep(500 * time.Millisecond)

		if i == 5 {
			cancelFunc()
		}
	}

}
