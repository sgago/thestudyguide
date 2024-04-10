package twopointers

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Two sum sorted
https://algo.monster/problems/two_sum_sorted

Given an array of integers sorted in ascending order,
find two numbers that add up to a given target.
Return the indices of the two numbers in ascending order.
You can assume elements in the array are unique and there is
only one solution. Do this in O(n) time and with constant auxiliary space.

Input:
	arr: a sorted integer array
target: the target sum we want to reach

Sample Input: [2 3 4 5 8 11 18], 8
Sample Output: 1 3

Target = 8

  l
2 3 4 5 8 11 18
      r

Shoudl work, l+r must sum to target.

It's sorted, so if l + r > target, pull in right side
if l+r < target, pull in left side.

target = l + r
target - left = right
target - right = left

Assumption array is sorted and there's at least one soln!

For sorted, we can do this in N with two pointers. Yay.
For unsorted, we can't do a bin search or anything.
Probably best to just sort and do this for NlogN + N => NlogN.
*/

func TestTwoSumSorted(t *testing.T) {
	nums := []int{2, 3, 4, 5, 8, 11, 18}
	slices.Sort(nums)

	l, r := twoSumSorted(nums, 8)

	fmt.Println("l r:", l, r)

	assert.Equal(t, 1, l)
	assert.Equal(t, 3, r)
}

func TestTwoSumSorted_TargetDoesNotExist(t *testing.T) {
	nums := []int{2, 3, 4, 5, 8, 11, 18}
	slices.Sort(nums)

	l, r := twoSumSorted(nums, 4)

	fmt.Println("l r:", l, r)

	assert.Equal(t, -1, l)
	assert.Equal(t, -1, r)
}

func twoSumSorted(nums []int, target int) (left int, right int) {
	if len(nums) == 0 {
		return -1, -1
	}

	right = len(nums) - 1
	curr := nums[left] + nums[right]

	for curr != target && left < right {
		if curr > target {
			right--
		} else {
			left++
		}

		curr = nums[left] + nums[right]
	}

	if left >= right {
		return -1, -1
	}

	return left, right
}
