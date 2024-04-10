package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Given a sorted array of integers and a target integer,
find the first occurrence of the target and return its index.
Return -1 if the target is not in the array.

Input:
	arr = [1, 3, 3, 3, 3, 6, 10, 10, 10, 100]
	target = 3
	Output: 1

Explanation: The first occurrence of 3 is at index 1.

Input:
	arr = [2, 3, 5, 7, 11, 13, 17, 19]
	target = 6
	Output: -1

Explanation: 6 does not exist in the array.
*/

func TestFirstOccurenceOf3(t *testing.T) {
	arr := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	target := 3

	actual := firstOccurence(arr, target)

	assert.Equal(t, 1, actual)
}

func TestFirstOccurenceOfDoesNotExist(t *testing.T) {
	arr := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	target := 7

	actual := firstOccurence(arr, target)

	assert.Equal(t, -1, actual)
}

func TestFirstOccurenceFirstElement(t *testing.T) {
	arr := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	target := 1

	actual := firstOccurence(arr, target)

	assert.Equal(t, 0, actual)
}

func TestFirstOccurenceLastElement(t *testing.T) {
	arr := []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}
	target := 100

	actual := firstOccurence(arr, target)

	assert.Equal(t, 9, actual)
}

func firstOccurence(arr []int, target int) int {
	ans, mid, left, right := -1, 0, 0, len(arr)-1

	// Start:
	// l              m
	// 1, 3, 3, 3, 3, 6, 10, 10, 10, 100
	//								 r

	// First pass
	// l     m
	// 1, 3, 3, 3, 3, 6, 10, 10, 10, 100
	//			   r

	// Second pass
	// lm
	// 1, 3, 3, 3, 3, 6, 10, 10, 10, 100
	//	  r

	// Third pass
	//    lm
	// 1, 3, 3, 3, 3, 6, 10, 10, 10, 100
	//	  r

	for left <= right {
		mid = left + (right-left)/2

		// This is the magic. We look while >= to the target, which'll keep the loop going
		// to find the first occurence. So, be sure to track the answer somewhere when doing binary search.
		if arr[mid] == target {
			ans = mid // This tracks the answer separately, cause we're moving mid, left, and right around
			right = mid - 1
		} else if arr[mid] > target {
			// Pull in the right side
			right = mid - 1
		} else {
			// Pull in the left side
			left = mid + 1
		}
	}

	return ans
}
