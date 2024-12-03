package linked_list

import (
	asserton "github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_New(t *testing.T) {
	assert := asserton.New(t)

	list := &LinkedList[int]{}

	assert.Equal(&LinkedList[int]{}, list, "Should initialize the linked list")
	assert.Equal((*Node[int])(nil), list.Head, "Should have a nil value for the head")
	assert.Equal((*Node[int])(nil), list.Tail, "Should have a nil value for the tail")
}

func TestLinkedList_GetHead(t *testing.T) {
	assert := asserton.New(t)

	list := &LinkedList[int]{}

	head := list.GetHead()

	assert.Equal((*Node[int])(nil), head, "Should be a node")
}
