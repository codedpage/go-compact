package abc_test

import (
	"abc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestABCImpl_DoSomething(t *testing.T) {
	obj := &abc.ABCImpl{}
	result := obj.DoSomething(10)
	assert.Equal(t, "Received 10", result)
}

func TestABCImpl_DoSomething2(t *testing.T) {
	obj := &abc.ABCImpl{}
	res, err := obj.DoSomething2("Manoj")
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Manoj", res)

	_, err = obj.DoSomething2("")
	assert.Error(t, err)
}
