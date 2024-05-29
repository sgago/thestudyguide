package getequalsubstringswithinbudget

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1208. Get Equal Substrings Within Budget
https://leetcode.com/problems/get-equal-substrings-within-budget

You are given two strings s and t of the same length and an integer maxCost.

You want to change s to t. Changing the ith character of s to ith character of
t costs |s[i] - t[i]| (i.e., the absolute difference between the ASCII values of the characters).

Return the maximum length of a substring of s that can be changed to be the same as
the corresponding substring of t with a cost less than or equal to maxCost.
If there is no substring from s that can be changed to its corresponding substring
from t, return 0.

Example 1:
	Input: s = "abcd", t = "bcdf", maxCost = 3
	Output: 3
	Explanation: "abc" of s can change to "bcd".
	That costs 3, so the maximum length is 3.

Example 2:
	Input: s = "abcd", t = "cdef", maxCost = 3
	Output: 1
	Explanation: Each character in s costs 2 to change to character in t,  so the maximum length is 1.

Example 3:
	Input: s = "abcd", t = "acde", maxCost = 0
	Output: 1
	Explanation: You cannot make any change, so the maximum length is 1.

Constraints:
- 1 <= s.length <= 105
- t.length == s.length
- 0 <= maxCost <= 106
- s and t consist of only lowercase English letters.
*/

/*
We have a budget (maxCost) for changing individual chars in s to t.
We want to get s as close to t as possible.
Cost is the abs diff between ascii chars like |s[i]-t[i]|.
Ah, ok, we want to return the max len of a substring that can be changed.
*/

func TestCheckStuff(t *testing.T) {
	fmt.Println('z' - 'a')
}

func TestGetEqualSubstringsWithinBudget_abcd_to_cdef(t *testing.T) {
	actual := getEqualSubstringsWithinBudget("abcd", "cdef", 3)

	fmt.Println(actual)

	assert.Equal(t, 1, actual)
}

func TestGetEqualSubstringsWithinBudget_aaa_to_zzz(t *testing.T) {
	actual := getEqualSubstringsWithinBudget("aaa", "zzz", 3)

	fmt.Println(actual)

	assert.Equal(t, 0, actual)
}

func TestGetEqualSubstringsWithinBudget_aa_to_aa(t *testing.T) {
	actual := getEqualSubstringsWithinBudget("aa", "aa", 0)

	fmt.Println(actual)

	assert.Equal(t, 2, actual)
}

func TestGetEqualSubstringsWithinBudget_abcd_to_bcdf(t *testing.T) {
	actual := getEqualSubstringsWithinBudget("abcd", "bcdf", 3)

	fmt.Println(actual)

	assert.Equal(t, 3, actual)
}

func getEqualSubstringsWithinBudget(s, t string, maxCost int) int {
	n := len(s)

	if n != len(t) {
		return -1 // Error, strings must be of equal length
	}

	maxCost = max(maxCost, 0)

	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = int(math.Abs(float64(s[i] - t[i])))
	}

	ans, curr, start := 0, 0, 0

	// Sliding window to find the longest substring within budget
	for end := 0; end < n; end++ {
		curr += costs[end]

		for curr > maxCost {
			curr -= costs[start]
			start++
		}

		ans = max(ans, end-start+1)
	}

	return ans
}
