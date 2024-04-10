package bubble

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arr := []int{5, 3, 1, 2, 4}
	Sort(arr)
	assert.True(t, slices.IsSorted(arr))
}

func TestSort_WithEmptyCollection(t *testing.T) {
	arr := []int{}
	Sort(arr)
	assert.True(t, slices.IsSorted(arr))
}
