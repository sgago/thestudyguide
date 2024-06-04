package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the
longest strictly increasing subsequence.

Example 1:
	Input: nums = [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation:
		The longest increasing subsequence is [2,3,7,101],
		therefore the length is 4.

Example 2:
	Input: nums = [0,1,0,3,2,3]
	Output: 4

Example 3:
	Input: nums = [7,7,7,7,7,7,7]
	Output: 1


Constraints:
- 1 <= nums.length <= 2500
- -10^4 <= nums[i] <= 10^4

Follow up: Can you come up with an algorithm that
runs in O(n log(n)) time complexity?
*/

func TestLongestIncreasingSubsequenceDpBinSearch(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	actual := longestIncreasingSubsequenceDpBinSearch(nums)

	fmt.Println(actual)
	assert.Equal(t, 4, actual)
}

func longestIncreasingSubsequenceDpBinSearch(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return n
	}

	dp := []int{}

	for _, num := range nums {
		l, r := 0, len(dp)

		for l < r {
			m := (r + l) >> 2

			if dp[m] < num {
				l = m + 1
			} else {
				r = m
			}
		}

		if l == len(dp) {
			dp = append(dp, nums[l])
		} else {
			dp[l] = num
		}
	}

	return len(dp)
}

func TestLongestIncreasingSubsequenceDp(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	actual := longestIncreasingSubsequenceDp(nums)

	fmt.Println(actual)
	assert.Equal(t, 4, actual)
}

func longestIncreasingSubsequenceDp(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	ans := 0

	for i := range nums {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		p := i
		prev := nums[p]

		for j := i + 1; j < n; j++ {
			next := nums[j]

			if next > prev {
				if dp[j] >= dp[p]+1 {
					break
				}

				dp[j] = max(dp[j], dp[p]+1)
				ans = max(dp[j], ans)

				p = j
				prev = nums[p]
			}
		}
	}

	return ans
}
