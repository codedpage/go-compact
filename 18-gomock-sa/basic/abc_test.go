package abc_test

import (
	"testing"

	// replace with your actual module path

	mocks "abc/mocks"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestDoSomething(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockABC := mocks.NewMockABCInterface(ctrl)
	mockABC.EXPECT().DoSomething(10).Return("Done")

	result := mockABC.DoSomething(10)
	assert.Equal(t, "Done", result)
}

func TestDoSomething2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockABC := mocks.NewMockABCInterface(ctrl)
	mockABC.EXPECT().DoSomething2("aa").Return("output", nil)

	result, err := mockABC.DoSomething2("aa")
	assert.NoError(t, err)
	assert.Equal(t, "output", result)

}
