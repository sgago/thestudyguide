package insertion

import (
	"cmp"

	sliceutil "sgago/thestudyguide/utils/slices"
)

// Sort a collection via insertion sort.
func Sort[T cmp.Ordered](arr []T) {
	SortFunc[T](arr, func(i, j T) bool { return i < j })
}

func SortFunc[T any](arr []T, less func(i, j T) bool) {
	for i := 1; i < len(arr); i++ {
		// This is the magic. j starts at the ith element and walks smaller elements to the front of the array.
		// So, while j-1 > j, swap it so that j-1 < j and decrement j. Note j > 0 cause we're doing j-1 stuff.
		for j := i; j > 0 && !less(arr[j-1], arr[j]); j-- {
			sliceutil.Swap(arr, j-1, j)
		}
	}
}
