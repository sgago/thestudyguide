package bubble

import (
	"cmp"

	sliceutil "sgago/thestudyguide/utils/slices"
)

// Sort a collection via bubble sort.
//
//   - Time complexity is O(N^2). Two loops. Worse than most other basic and advanced sorts.
//   - Stable: Yes. Does not swap elements with equal values.
//   - In-place: Yes. Only requires O(1) additional memory.
//   - Adaptive: Yes. Efficient for mostly sorted input collections. Only O(KN) when input is K places away.
//
// Example:
//
//	[3 5 1 2 4] - Swap(3, 5)
//	[3 1 5 2 4] - Swap(1, 5)
//	[3 1 2 5 4] - Swap(2, 5)
//	[3 1 2 4 5] - Swap(4, 5)
//	  Swapped: true
//	[1 3 2 4 5] - Swap(1, 3)
//	[1 2 3 4 5] - Swap(2, 3)
//	  Swapped: true
//	  Swapped: false
//	  Done
func Sort[T cmp.Ordered](arr []T) {
	SortFunc[T](arr, func(i, j T) bool { return i < j })
}

func SortFunc[T any](arr []T, less func(i, j T) bool) {
	for {
		swapped := false

		for i := 1; i < len(arr); i++ {
			if !less(arr[i-1], arr[i]) {
				sliceutil.Swap(arr, i-1, i)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
