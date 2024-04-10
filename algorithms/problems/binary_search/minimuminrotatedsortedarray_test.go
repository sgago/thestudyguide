package binarysearch

import "testing"

/*
Find Minimum in Rotated Sorted Array
https://algo.monster/problems/min_in_rotated_sorted_array

A sorted array of unique integers was rotated at an unknown pivot. For example,
[10, 20, 30, 40, 50] becomes [30, 40, 50, 10, 20].
Find the index of the minimum element in this array.
	Input: [30, 40, 50, 10, 20]
	Output: 3
Explanation: the smallest element is 10 and its index is 3.

	Input: [3, 5, 7, 11, 13, 17, 19, 2]
	Output: 7
Explanation: the smallest element is 2 and its index is 7.

We want the min number, but the sorted array is rotated.

	                 lm
	[30, 40, 50, 10, 20]
				 r

	for left <= right?
	if mid > r
		pull in left, solution is on the right side
	if mid < l
		pull in right, solution is over on the left

Values are unique
*/

func TestMinInRotatedSortedArray(t *testing.T) {
	//findMinInRotatedSortedArray([]int{30, 40, 50, 10, 20})

	findMinInRotatedSortedArray([]int{3, 5, 7, 11, 13, 17, 19, 2})
}

func findMinInRotatedSortedArray(arr []int) int {
	ans := -1
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2 // FIXME

		if arr[mid] > arr[right] {
			left = mid + 1
		} else if arr[mid] <= arr[right] {
			ans = mid
			right = mid - 1
		}
	}

	return ans
}
