package numberofsubarrayswithoddsum

import (
	"testing"
)

/*
1524. Number of Sub-arrays With Odd Sum
https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/

Given an array of integers arr, return the number of subarrays
with an odd sum.

Since the answer can be very large, return it modulo 10^9 + 7.

Example 1:
	Input: arr = [1,3,5]
	Output: 4
	Explanation:
		All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
		All sub-arrays sum are [1,4,9,3,8,5].
		Odd sums are [1,9,3,5] so the answer is 4.

Example 2:
	Input: arr = [2,4,6]
	Output: 0
	Explanation:
		All subarrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
		All sub-arrays sum are [2,6,12,4,10,6].
		All sub-arrays have even sum and the answer is 0.

Example 3:
	Input: arr = [1,2,3,4,5,6,7]
	Output: 16

Constraints:
- 1 <= arr.length <= 105
- 1 <= arr[i] <= 100

================================================================

Odd and even addition rules are:
Odd + Odd = Even = 1 + 1 = 0
Odd + Even = Odd = 1 + 0 = 1
Even + Even = Even = 0 + 0 = 0

This is like XOR on the LSB.

1 3 5 is odd odd odd is 1, 3, 5, and 9 as answers.
The prefix sum is 1 4 9. Does this help?
These are overlapping subproblems. Does DP help?
We can also generate every subarray and check if the sum is odd.

Init:
dp[0][0] = 1
dp[1][1] = 1
dp[2][2] = 1

dp[0][1] = dp[0] ^ dp[1] = 0
dp[1][2] = dp[1] ^ dp[2] = 0

dp[0][2] = dp[0] ^ dp[2] = 1

If right idx is max, then quit; otherwise, continue.

Another way to solve this monster is to DFS out all the combos and add memoizing.
Watch out, they need to be contiguous to be a subarray. The keyword is subarray.

   0  1  2 ith
0 00 01 02
1 xx 11 12
2 xx xx 22
jth

Also, because we can can skip using the sums and only look at true or false,
I can use a bitmasks to represent the odd/even sums.
*/

func TestNumOfSubarrayWithOddSum(t *testing.T) {

}
