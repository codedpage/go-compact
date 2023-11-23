package main

import "testing"

func TestSum(t *testing.T) {

	n := 2
	if n != 2 {
		t.Error("Expected 1 + 1 =  2 but got", n)
	}
}

// go test sum_test.go
// --- FAIL: TestSum (0.00s)
//     sum_test.go:20: Expected 1 + 1 =  2 but got 1
