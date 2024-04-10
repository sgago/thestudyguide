package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive_WithZeroAndNegatives(t *testing.T) {
	nums := []int{0, -2, 1, -1, 3} // Ans: 2 is the first missing positive num.
	actual := firstMissingPositiveNumber(nums)
	assert.Equal(t, 2, actual)
}

func TestFirstMissingPositive_WithNoNegativesOrZero(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2} // Ans: 2 is the first missing positive num.
	actual := firstMissingPositiveNumber(nums)
	assert.Equal(t, 6, actual)
}

func firstMissingPositiveNumber(nums []int) int {
	low := math.MaxInt
	high := 0
	good := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			continue
		}

		low = min(low, nums[i])
		high = max(high, nums[i])

		for nums[i] != i {
			if nums[i] < 0 || nums[i] > len(nums)-1 {
				break
			}

			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]

			if i != 0 && nums[i] == nums[i-1]+1 {
				good++
			}
		}
	}

	return -1
}
