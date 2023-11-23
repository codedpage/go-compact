package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Singal breaker (Cancel)
func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason

	time.Sleep(100 * time.Millisecond)
	return errors.New("operation1 failed") //fmt.Errorf("operation1 failed")
	//return nil
}

// main work
func operation2(ctx context.Context) {

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println(ctx, "3..............operation2 done")
	case <-ctx.Done():
		fmt.Println("operation2 halted because operation1 returns error")
	}
}

// If this operation returns an error cancel all operations using this context
func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())
	fmt.Println(ctx, "1----------", cancelFunc)

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		fmt.Println(ctx, "2============", err)

		if err != nil {
			cancelFunc()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
