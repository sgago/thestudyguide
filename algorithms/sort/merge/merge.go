package merge

import "cmp"

func Sort[T cmp.Ordered](arr *[]T) {
	SortFunc[T](arr, func(i, j T) bool { return i < j })
}

func SortFunc[T any](arr *[]T, less func(i, j T) bool) {
	sortFunc[T](arr, less)
}

// sortFunc performs a merge sort on the given slice using the provided comparison function
// It modifies the original slice and returns a reference to it
func sortFunc[T any](arrP *[]T, less func(i, j T) bool) *[]T {
	arr := *arrP

	// Base case: If the slice has 0 or 1 elements, it's already sorted
	if len(arr) <= 1 {
		return arrP
	}

	// Split the slice into two halves
	mid := len(arr) / 2
	lArr, rArr := arr[:mid], arr[mid:]

	// Recursively sort the left and right halves
	sortFunc[T](&lArr, less)
	sortFunc[T](&rArr, less)

	// Merge the sorted halves back into the original slice
	i, j, k := 0, 0, 0
	for i < len(lArr) && j < len(rArr) {
		if less(lArr[i], rArr[j]) {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
		k++
	}

	// Copy any remaining elements from the left or right slice
	copy(arr[k:], lArr[i:])
	copy(arr[k:], rArr[j:])

	return arrP
}
