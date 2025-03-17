package myutil

import (
	"fmt"
	"testing"
)

func TestAddMock(t *testing.T) {
	result := Add(4, 3)

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	expected := 7
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}

}

func TestSubtractMock(t *testing.T) {
	result := Subtract(6, 9)
	expected := -3
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}
}

func TestConcatenateMock(t *testing.T) {
	result := Concatenate("Learning", "Go")
	expected := "Learning Go"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}

func TestGreetMock(t *testing.T) {
	result := Greet("Ayesha")
	expected := "Hello Ayesha!"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}
