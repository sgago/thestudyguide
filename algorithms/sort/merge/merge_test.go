package merge

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortFunc(t *testing.T) {
	arr := []int{5, 3, 1, 2, 4}

	SortFunc[int](&arr, func(i, j int) bool {
		return i < j
	})

	assert.True(t, slices.IsSorted(arr))
}
