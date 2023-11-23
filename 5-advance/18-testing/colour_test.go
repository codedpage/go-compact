package main

import "testing"

func TestMsg(t *testing.T) {
	a := "red"

	if a != "green" {
		t.Error("Exepected green but got", a)
	}
}

func TestShape(t *testing.T) {
	a := "triangle"

	if a != "circle" {
		t.Error("Exepected circle but got", a)
	}
}

//go test colour_test.go
// --- FAIL: TestMsg (0.00s)
//     colour_test.go:9: Exepected green but got red
// --- FAIL: TestShape (0.00s)
//     colour_test.go:17: Exepected circle but got triangle
