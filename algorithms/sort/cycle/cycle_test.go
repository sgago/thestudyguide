package cycle

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortInt(t *testing.T) {
	nums := []int{3, 1, 2}
	Sort(nums)
	assert.True(t, slices.IsSorted(nums))
}

func TestSortInt_WithHigherNumber(t *testing.T) {
	nums := []int{6, 3, 4, 1, 7}
	Sort(nums)
	assert.True(t, slices.IsSorted(nums))
}
