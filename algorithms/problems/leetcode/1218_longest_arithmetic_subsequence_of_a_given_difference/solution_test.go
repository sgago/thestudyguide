package longestarithmeticsubsequenceofagivendifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1218. Longest Arithmetic Subsequence of Given Difference
https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference

Given an integer array arr and an integer difference,
return the length of the longest subsequence in arr which is an arithmetic sequence
such that the difference between adjacent elements in the subsequence equals difference.

A subsequence is a sequence that can be derived from arr by deleting
some or no elements without changing the order of the remaining elements.

Example 1:
	Input: arr = [1,2,3,4], difference = 1
	Output: 4
	Explanation: The longest arithmetic subsequence is [1,2,3,4].

Example 2:
	Input: arr = [1,3,5,7], difference = 1
	Output: 1
	Explanation: The longest arithmetic subsequence is any single element.

Example 3:
	Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
	Output: 4
	Explanation: The longest arithmetic subsequence is [7,5,3,1].

Constraints:
- 1 <= arr.length <= 105
- -104 <= arr[i], difference <= 104
*/

/*
So, we get a bunch of nums (arr) and a difference.
We need to count the longest seq where the difference
between numbers exactly equals the difference.
1,2,3,4 and 1 -> 4
1,2,3,4 and 2 -> 2 (1,3 or 2,4)

We can brute force with multiple loops.
We could also sort (nlogn) and then count? No, that's ugly.
This is giving me longest common subsequence vibes. That's optimal via DP/DFS.

Can we do this with two pointer? Or is that just a brute force here?
Without a memo, we have to check like 10+9+8+7... which is like n**2, (loops are factors for tc)
I *think* dp is a bit cleaner here, we can stop looping if there's a better.
soln stored in our memo.

dp memo will hold seq lengths:
	1 5 7 8 5 3 4 2 1, k=-2
	1 1 1 1 1 1 1 1 1 = init
0   1 1 1 1 1 1 1 1 1 = idx 0 is just 1s
1     1 1 1 1 2 1 1 3 = idx 1 is 5 3 1 and is 3
2       1 1 2 3 1 1 4 = idx 2 is 7 5 3 1
3         1 2 3 1 1 4 = idx 3 is just 8 so 1
4           2           idx 4 is >1, we got here with a bigger seq, so just skip it
5             3         idx 5 is >1, we got here with a bigger seq, so skip it again, also curr max == 4
6               1       curr max is 4, there's no way this can be a soln, so skip it.
*/

func TestLongestArithmeticSubsequenceOfGivenDifference_1234_diff1(t *testing.T) {
	actual := longestSequence([]int{1, 2, 3, 4}, 1)

	assert.Equal(t, 4, actual)
}

func TestLongestArithmeticSubsequenceOfGivenDifference_1234_diff2(t *testing.T) {
	actual := longestSequence([]int{1, 2, 3, 4}, 2)

	assert.Equal(t, 2, actual)
}

func TestLongestArithmeticSubsequenceOfGivenDifference_NoDiffs(t *testing.T) {
	actual := longestSequence([]int{1, 3, 5, 7}, 1)

	assert.Equal(t, 1, actual)
}

func TestLongestArithmeticSubsequenceOfGivenDifference_Neg(t *testing.T) {
	actual := longestSequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2)

	assert.Equal(t, 4, actual)
}

func longestSequence(nums []int, diff int) int {
	n := len(nums)

	dp := make([]int, n)
	ans := 1

	for i := 0; i < n && ans < n-i; i++ {
		dp[i] = max(dp[i], 1)
		prev := i

		for j := i + 1; j < n; j++ {
			if nums[j]-nums[prev] == diff {
				if dp[j] > dp[prev]+1 {
					// This number is part of a longer seq already
					// Move to next number, there's no point in continuing
					// because this seq will just be shorter
					break
				}

				dp[j] = max(dp[j], dp[prev]+1)
				ans = max(ans, dp[j])
				prev = j
			}
		}
	}

	return ans
}
