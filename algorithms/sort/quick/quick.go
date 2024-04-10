package quick

import sliceutil "sgago/thestudyguide/utils/slices"

// SortFunc performs quicksort on the given slice using the provided comparison function
func SortFunc[T any](arr *[]T, less func(i, j T) bool) {
	sortFunc[T](arr, 0, len(*arr)-1, less)
}

// sortFunc is a helper function for SortFunc that implements the quicksort algorithm
func sortFunc[T any](arr *[]T, low, high int, less func(i, j T) bool) {
	if low < high {
		// Partition the array and get the index of the pivot element
		idx := partition(arr, low, high, less)
		// Recursively sort the sub-arrays before and after the pivot
		sortFunc(arr, low, idx-1, less)
		sortFunc(arr, idx+1, high, less)
	}
}

// partition is a helper function for sortFunc that partitions the array around a pivot element
func partition[T any](arr *[]T, low, high int, less func(i, j T) bool) int {
	a := *arr // Dereference for easy use

	pivot := a[high] // Choose the last element as our pivot
	i := low - 1     // Last index of our elements that are <= the pivot, it is the slow pointer

	// Iterate through the array and rearrange elements based on the pivot
	// j is the last index of our elements that are > the pivot, it is the fast pointer
	for j := low; j < high; j++ {
		if !less(pivot, a[j]) {
			i++
			// Swap elements using the provided Swap function
			sliceutil.Swap(a, i, j)
		}
	}

	// Swap the pivot element to its correct position
	sliceutil.Swap(a, i+1, high)
	return i + 1
}
