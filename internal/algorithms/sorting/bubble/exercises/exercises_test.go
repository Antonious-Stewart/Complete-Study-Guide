package exercises

import (
	assertion "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"slices"
	"testing"
)

type Spy struct {
	mock.Mock
}

func (sp *Spy) isSorted(s []int) bool {
	args := sp.Called(s)
	return args.Bool(0)
}

func TestSortInAscendingOrder(t *testing.T) {
	input := []int{56, 2324, 34, 6, 23, 2, 566, 8, 3}
	assert := assertion.New(t)
	spy := Spy{}

	spy.On("isSorted", input).Return(false)
	SortInAscendingOrder(input, spy.isSorted)

	spy.AssertCalled(t, "isSorted", input)
	spy.AssertNumberOfCalls(t, "isSorted", 1)
	assert.True(slices.Equal([]int{2, 3, 6, 8, 23, 34, 56, 566, 2324}, input))
}

func TestSortInAscendingOrderEdgeCase_1(t *testing.T) {
	input := []int{3}
	assert := assertion.New(t)
	spy := Spy{}

	spy.On("isSorted", input).Return(true)
	SortInAscendingOrder(input, spy.isSorted)

	spy.AssertCalled(t, "isSorted", input)
	spy.AssertNumberOfCalls(t, "isSorted", 1)
	assert.True(spy.isSorted(input))
	assert.True(slices.Equal([]int{3}, input))
}

func TestSortInAscendingOrderEdgeCase_2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert := assertion.New(t)
	spy := Spy{}

	spy.On("isSorted", input).Return(true)
	SortInAscendingOrder(input, spy.isSorted)

	spy.AssertCalled(t, "isSorted", input)
	spy.AssertNumberOfCalls(t, "isSorted", 1)
	assert.True(spy.isSorted(input))
}

func TestSortInAscendingOrderEdgeCase_3(t *testing.T) {
	input := []int{4, 2, 2, 1}
	assert := assertion.New(t)
	spy := Spy{}

	spy.On("isSorted", input).Return(false)
	SortInAscendingOrder(input, spy.isSorted)

	spy.AssertCalled(t, "isSorted", input)
	spy.AssertNumberOfCalls(t, "isSorted", 1)
	assert.False(spy.isSorted(input))
	assert.True(slices.Equal([]int{1, 2, 2, 4}, input))
}
