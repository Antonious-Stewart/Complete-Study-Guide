package queue

import (
	"fmt"
	assertion "github.com/stretchr/testify/assert"
	"testing"
)

var queue *Queue = New()

func TestNewQueue(t *testing.T) {
	assert := assertion.New(t)

	fmt.Print(&queue.Data)
	assert.Nil(
		[]interface{}(nil),
		queue.Data,
		"should have nil data for the array",
	)
	assert.Equal(
		0,
		queue.Size(),
		"should be zero upon initialization",
	)
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assertion.New(t)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	assert.Len(queue.Data, 4, "Should have a value of 4")
	assert.Equal(4, queue.Size(), "length should update to be 4")
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assertion.New(t)

	value := queue.Dequeue()

	assert.Equal(1, value, "Should have the value of one")
	assert.NotEmpty(queue.Data, "Should not be empty")
	assert.Equal([]interface{}{2, 3, 4}, queue.Data, "Should equal the expected array")
	assert.Equal(3, queue.Size(), "should have a new value for the size")
}

func TestQueue_Front(t *testing.T) {
	assert := assertion.New(t)

	front := queue.Front()

	assert.Equal(2, front, "Should be the front of the queue")
	assert.Len(queue.Data, 3, "should still have the length of 3")
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assertion.New(t)

	assert.False(queue.IsEmpty(), "should not be empty")

	newQueue := &Queue{
		Data:   nil,
		length: 0,
	}

	assert.True(newQueue.IsEmpty(), "should be true with a new queue")
	assert.Equal(0, newQueue.Size(), "should be 0")
}
