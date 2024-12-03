package stack

import (
	assertion "github.com/stretchr/testify/assert"
	"testing"
)

var stack = Stack[int]{
	Data: []int{},
}

func TestPushMethod(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal([]interface{}{}, stack.Data, "should be nil")

	stack.Push(1)
	assert.Equal(1, stack.Data[0], "should be 1")
}

func TestPopMethod(t *testing.T) {
	assert := assertion.New(t)

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	value, _ := stack.Pop()

	assert.Equal(4, value, "should be 4")
}

func TestPeekMethod(t *testing.T) {
	assert := assertion.New(t)

	value, _ := stack.Peek()
	assert.Equal(3, value, "should be 3")
}

func TestIsEmptyMethod(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(false, stack.IsEmpty(), "should be false")

	newStack := Stack[int]{}

	assert.Equal(true, newStack.IsEmpty(), "should be true")
}

func TestSizeMethod(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(3, stack.Size(), "should be 3")
}
