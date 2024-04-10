/*
FirstElementNotSmallerThanTarget
https://algo.monster/problems/binary_search_first_element_not_smaller_than_target

Given an array of integers sorted in increasing order and a target,
find the index of the first element in the array that is larger than or equal to the target.
Assume that it is guaranteed to find a satisfying number.

Input:
			  |
			  v
	arr = [1, 3, 3, 5, 8, 8, 10]
	target = 2
	Output: 1

Explanation: the first element larger than 2 is 3 which has index 1.

Input:
					|
					v
	arr = [2, 3, 5, 7, 11, 13, 17, 19]
	target = 6
	Output: 3

Explanation: the first element larger than 6 is 7 which has index 3.
*/

package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FirstElementNotSmallerThanTarget returns the index of the first element not smaller
// than a given target in the input array. It is assumed that there is a satisifying solution.
// An empty slice returns -1.
func FirstElementNotSmallerThanTarget(target int, vals ...int) int {
	if len(vals) == 0 {
		return -1
	} else if len(vals) == 1 {
		return 0
	}

	left, right := 0, len(vals)-1

	// There are like two types of binary searches:
	// 1. We find the exact target we're looking for
	// 2. There's no "precise exact target" and we need to look at a boundary
	// This is boundary one.
	for left <= right {
		mid := left + (right-left)/2

		if vals[mid] < target {
			// mid is >= target, so pull in the left side, soln is to the right of mid
			// When doing a boundary condition, like this one, left will conviently hold our solution
			left = mid + 1
		} else {
			// mid is >= target, so pull in the right side, soln is to the left of mid
			right = mid - 1
		}
	}

	return left
}

// Array is monotonic increasing
// Find index of the first element in the array that is larger than or equal to the target.
// Assume you'll find a solution.
// Looks like a binary search will work here, we've got a sorted collection and can check for a boundary condition.
func TestFirstElementNotSmallerThanTarget_Example1(t *testing.T) {
	target := 2
	expected := 1
	actual := FirstElementNotSmallerThanTarget(target, 1, 3, 3, 5, 8, 8, 10)

	assert.Equal(t, expected, actual)
}

func TestFirstElementNotSmallerThanTarget_Example2(t *testing.T) {
	target := 6
	expected := 3
	actual := FirstElementNotSmallerThanTarget(target, 2, 3, 5, 7, 11, 13, 17, 19)

	assert.Equal(t, expected, actual)
}

func TestFirstElementNotSmallerThanTarget_FirstElement(t *testing.T) {
	target := 1
	expected := 0
	actual := FirstElementNotSmallerThanTarget(target, 1, 3, 3, 5, 8, 8, 10)

	assert.Equal(t, expected, actual)
}

func TestFirstElementNotSmallerThanTarget_LastElement(t *testing.T) {
	target := 10
	expected := 6
	actual := FirstElementNotSmallerThanTarget(target, 1, 3, 3, 5, 8, 8, 10)

	assert.Equal(t, expected, actual)
}

func TestFirstElementNotSmallerThanTarget_SmallArrayFirstElement(t *testing.T) {
	target := 1
	expected := 0
	actual := FirstElementNotSmallerThanTarget(target, 1, 3)

	assert.Equal(t, expected, actual)
}

func TestFirstElementNotSmallerThanTarget_SmallArrayLastElement(t *testing.T) {
	target := 3
	expected := 1
	actual := FirstElementNotSmallerThanTarget(target, 1, 3)

	assert.Equal(t, expected, actual)
}
