package bubble

import (
	assertion "github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	assert := assertion.New(t)
	input := []int{45, 3, 87, 12, 99, 34, 76, 8, 23, 56, 91, 67}
	expected := []int{3, 8, 12, 23, 34, 45, 56, 67, 76, 87, 91, 99}

	Sort(input)

	assert.Len(input, len(expected), "should keep the same length for the array")
	assert.Equal(expected, input, "should sort the array")
}
