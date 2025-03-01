package lengthoflongestfibonaccisubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
873. Length of Longest Fibonacci Subsequence
https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/

A sequence x1, x2, ..., xn is Fibonacci-like if:
- n >= 3
- xi + xi+1 == xi+2 for all i + 2 <= n

Given a strictly increasing array arr of positive integers forming a sequence,
return the length of the longest Fibonacci-like subsequence of arr.
If one does not exist, return 0.

A subsequence is derived from another sequence arr by deleting any number of elements
(including none) from arr, without changing the order of the remaining elements.
For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].

Example 1:
	Input: arr = [1,2,3,4,5,6,7,8]
	Output: 5
	Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].

Example 2:
	Input: arr = [1,3,7,11,12,14,18]
	Output: 3
	Explanation:
		The longest subsequence that is fibonacci-like: [1,11,12], [3,11,14] or [7,11,18].

Constraints:
- 3 <= arr.length <= 1000
- 1 <= arr[i] < arr[i + 1] <= 109

==========================================================================================

This looks like a dynamic programming problem.
The output is longest length, so our dp[i] memo will hold lengths.
Each number by itself is a fibonacci sequence of length 1.
We'll init all dp elements to 1.

So, dp[i] = dp[prev - 1] + 1, but only when dp[prev - 1] + dp[prev - 2] == arr[i].
So, from ith element, we'll walk prev back all the way to prev - 2 == 0.

Ehhh, but we can drop numbers. This means we need to track the last two indices in our memo.
So this is a another 2D dp memo problem.

*/

func TestLengthOfLongestFibonacciSequence(t *testing.T) {
	actual := LongestFib([]int{1, 2, 3, 4, 5, 6, 7, 8})
	assert.Equal(t, 5, actual)
}

func LongestFib(arr []int) int {
	longest := 1

	return longest
}
