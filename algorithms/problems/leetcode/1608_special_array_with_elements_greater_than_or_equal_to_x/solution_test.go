package specialarraywithelementsgreaterthanorequaltox

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1608. Special Array With X Elements Greater Than or Equal X
https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x

You are given an array nums of non-negative integers.
nums is considered special if there exists a number x
such that there are exactly x numbers in nums that are
greater than or equal to x.

Notice that x does not have to be an element in nums.

Return x if the array is special, otherwise, return -1.
It can be proven that if nums is special, the value for x is unique.

Example 1:
	Input: nums = [3,5]
	Output: 2
	Explanation:
		There are 2 values (3 and 5) that are greater
		than or equal to 2.

Example 2:
	Input: nums = [0,0]
	Output: -1
	Explanation:
		No numbers fit the criteria for x.
		If x = 0, there should be 0 numbers >= x, but there are 2.
		If x = 1, there should be 1 number >= x, but there are 0.
		If x = 2, there should be 2 numbers >= x, but there are 0.
		x cannot be greater since there are only 2 numbers in nums.

Example 3:
	Input: nums = [0,4,3,0,4]
	Output: 3
	Explanation:
		There are 3 values that are greater than or equal to 3.

Constraints:
- 1 <= nums.length <= 100
- 0 <= nums[i] <= 1000
*/

/*
Mmk, so we're trying to find a num x in nums
that is has x many >= numbers. Nums are all positive
and can be zero. x does NOT have to be an element
in nums.

The formula is like saying there are "x values >= to x".

We could sort nums (nlogn) and then loop thru
(n maybe logn with bin search) to find a number that exists.

Can we do this without sorting though?
*/

func TestSpecialArrayGreaterThanOrEqualToX_35(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{3, 5})

	fmt.Println(actual)
	assert.Equal(t, 2, actual)
}

func TestSpecialArrayGreaterThanOrEqualToX_Empty(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{})

	fmt.Println(actual)
	assert.Equal(t, 0, actual)
}

func TestSpecialArrayGreaterThanOrEqualToX_1(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{1})

	fmt.Println(actual)
	assert.Equal(t, 1, actual)
}

func TestSpecialArrayGreaterThanOrEqualToX_00(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{0, 0})

	fmt.Println(actual)
	assert.Equal(t, -1, actual)
}

func TestSpecialArrayGreaterThanOrEqualToX_04304(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{0, 4, 3, 0, 4})

	fmt.Println(actual)
	assert.Equal(t, 3, actual)
}

func TestSpecialArrayGreaterThanOrEqualToX_05678(t *testing.T) {
	actual := specialArrayGreaterThanOrEqualToX([]int{0, 5, 6, 7, 8})

	fmt.Println(actual)
	assert.Equal(t, 4, actual)
}

func specialArrayGreaterThanOrEqualToX(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] > 0 {
			return 1
		}
		return -1
	}

	slices.Sort(nums)

	n := len(nums)
	l, r := 0, n-1

	// Trying each number n, logn times
	// Feels somewhat novel and/or inefficient
	for x := 2; x <= n; x++ {
		for l < r {
			m := (r - l + l) >> 1
			mid := nums[m]

			if mid >= x {
				r = m
			} else {
				l = m + 1
			}
		}

		countGreaterThanOrEqualToX := n - l

		if countGreaterThanOrEqualToX == x {
			return x
		}
	}

	return -1
}
