package twopointers

import (
	"fmt"
	"slices"
	"testing"
)

/*
Move Zeroes
https://algo.monster/problems/move_zeros

Given an array of integers, move all the 0s to
the back of the array while maintaining the relative order
of the non-zero elements.

Do this in-place using constant auxiliary space.

Input: [1, 0, 2, 0, 0, 7]
Output: [1, 2, 7, 0, 0, 0]
*/

func TestMoveZeroes(t *testing.T) {
	actual := moveZeroes([]int{1, 0, 2, 0, 0, 7})

	fmt.Println(actual)
}

func moveZeroes(nums []int) []int {
	for len(nums) == 0 {
		return nums
	}

	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j := i

			for ; nums[j] == 0 && j < len(nums); j++ {
				cnt++
			}

			nums = slices.Delete(nums, i, j)
		}
	}

	nums = append(nums, make([]int, cnt)...)

	return nums
}
