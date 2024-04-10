package quick

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortFunc(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	SortFunc(&arr, func(i, j int) bool { return i < j })
	assert.True(t, slices.IsSorted(arr))
}

// func TestSortFunc(t *testing.T) {
// 	arr := []int{5, 3, 1, 4, 2, 7, 0, 4}
// 	fmt.Println("Unsorted array:", arr)

// 	quicksort(arr, 0, len(arr)-1)

// 	fmt.Println("Sorted array in place:", arr)
// }
