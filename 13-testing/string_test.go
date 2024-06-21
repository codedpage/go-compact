package myutil

import "testing"

func TestConcatenate(t *testing.T) {
	result := Concatenate("Learning", "Go")
	expected := "Learning Go"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}

func TestGreet(t *testing.T) {
	result := Greet("Golang")
	expected := "Hello Golang!"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}
