package myutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(4, 3)
	expected := 7

	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(6, 9)
	expected := -3
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}
}
