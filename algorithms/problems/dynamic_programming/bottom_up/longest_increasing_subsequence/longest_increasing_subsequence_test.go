package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/monotonic/stack"
)

/*
	Longest Increasing Subsequence (LIS)
	https://algo.monster/problems/longest_increasing_subsequence
*/

func TestLis(t *testing.T) {
	lis := LongestIncreasingSubsequence([]int{1, 2, 4, 3})
	fmt.Println(lis)
}

func LongestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums)+1) // Create memo, with +1 for init cond
	// dp[0] = 0 // It's zero by default, but our memo's init cond is 0
	length := 0

	for i := 1; i < len(nums); i++ {
		ni := nums[i-1]
		dp[i] = 1 // Any number is at least 1 or dp[0]+1==1

		for j := 1; j < i; j++ {
			nj := nums[j-1]

			if nj < ni {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		length = max(length, dp[i])
	}

	return length
}

// A nifty use of a monotonic stack to get NlogN TC.
func TestLisMonotonicStack(t *testing.T) {
	s := stack.New[int](10)

	length := 0

	for _, num := range []int{0, 1, 3, 2, 4, 5, -1, 0, 3} {
		length = max(length, len(s.Push(num))+s.Len()-1)
	}

	fmt.Println(length)
}
