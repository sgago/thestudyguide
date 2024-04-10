package selection

import (
	"cmp"

	sliceutil "sgago/thestudyguide/utils/slices"
)

// Sort a collection via selection sort.
//   - Time complexity is O(N^2). Two loops. Note that inner loop iterations is a factor of the outer.
//   - More efficient than bubble or selection sort; worse than the advanced quick, merge, or heap sorts.
//   - Stable: Yes. Does not swap elements with equal values.
//   - In-place: Yes, this version is, other selection sorts allocating a new array are not. Only requires O(1) additional memory. This one is more like insertion sort.
//   - Adaptive: Yes. Efficient for mostly sorted input collections. Only O(KN) when input is K places away.
//
// Example:
//
//	[5 3 1 2 4] - Swap(1, 5)
//	[1 3 5 2 4] - Swap(2, 3)
//	[1 2 5 3 4] - Swap(3, 5)
//	[1 2 3 5 4] - Swap(4, 5)
//	[1 2 3 4 5] - Done
func Sort[T cmp.Ordered](arr []T) {
	SortFunc[T](arr, func(i, j T) bool { return i < j })
}

func SortFunc[T any](arr []T, less func(i, j T) bool) {
	for i := 0; i < len(arr); i++ {
		minIdx := i

		for j := i + 1; j < len(arr); j++ {
			if less(arr[j], arr[minIdx]) {
				minIdx = j
			}
		}

		if minIdx != i {
			sliceutil.Swap(arr, minIdx, i)
		}
	}
}
