package twopointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Subarray sum and sliding window intro
	https://algo.monster/problems/subarray_sum_fixed

	Given an array (list) nums consisted of only non-negative integers,
	find the largest sum among all subarrays of length k in nums.

	For example, if the input is nums = [1, 2, 3, 7, 4, 1], k = 3,
	then the output would be 14 as the largest length 3 subarray sum
	is given by [3, 7, 4] which sums to 14.

	Slide along an array and find the largest sum. What's fun here, is
	we can add and remove values from the sum as we go along.
*/

func TestSubArraySum(t *testing.T) {
	nums := []int{1, 2, 3, 7, 4, 1}
	k := 3

	sum := subArrySum(nums, k)

	fmt.Println("sum:", sum)
	assert.Equal(t, 14, sum)
}

func TestSubArraySum_ArrayLenLessThanK(t *testing.T) {
	nums := []int{1, 2}
	k := 3

	sum := subArrySum(nums, k)

	fmt.Println("sum:", sum)
	assert.Equal(t, 3, sum)
}

func subArrySum(nums []int, k int) int {
	total := 0
	curr := 0

	for i := 0; i < len(nums) && i < k; i++ {
		curr += nums[i]
	}

	total = max(total, curr)

	if len(nums) <= k {
		return total
	}

	start, end := 0, k-1

	for end < len(nums)-1 {
		start++
		end++

		curr -= nums[start-1]
		curr += nums[end]

		total = max(total, curr)
	}

	return total
}
