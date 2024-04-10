package semi

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemi_Basic(t *testing.T) {
	nums := []int{4, 2, 0, 3, 1}
	actual := Sort(nums, 0)
	assert.True(t, slices.IsSorted(actual))
}

func TestSemi_NumsWithDifferentLowRange(t *testing.T) {
	nums := []int{4, 2, 0, 3, 1, 99, 8}
	actual := Sort(nums, 2)
	fmt.Println("nums:", actual)
}
