package myutil

import (
	"testing"
)

func TestCheck(t *testing.T) {

	l := []int{1, 2, 3, 4}
	//l := []int{1, 2, 3, 4, 5} //fail

	result := Check(l)
	if result != true {
		t.Errorf("Result was incorrect, got: %t", result)
	}

}
