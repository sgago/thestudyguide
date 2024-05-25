package longestidealsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
2370. Longest Ideal Subsequence
https://leetcode.com/problems/longest-ideal-subsequence

You are given a string s consisting of lowercase letters and an integer k.
We call a string t ideal if the following conditions are satisfied:
- t is a subsequence of the string s.
- The absolute difference in the alphabet order of every two adjacent letters in t is less than or equal to k.
Return the length of the longest ideal string.

A subsequence is a string that can be derived from another string by deleting some or
no characters without changing the order of the remaining characters.

Note that the alphabet order is not cyclic. For example, the absolute difference
in the alphabet order of 'a' and 'z' is 25, not 1.

Example 1:
	Input: s = "acfgbd", k = 2
	Output: 4
	Explanation: The longest ideal string is "acbd". The length of this string is 4, so 4 is returned.
	Note that "acfgbd" is not ideal because 'c' and 'f' have a difference of 3 in alphabet order.

Example 2:
	Input: s = "abcd", k = 3
	Output: 4
	Explanation: The longest ideal string is "abcd". The length of this string is 4, so 4 is returned.


Constraints:
- 1 <= s.length <= 105
- 0 <= k <= 25
- s consists of lowercase English letters.
*/

/*
So, basically, every letter in a string s needs to be within k distance of previous letter.
k == 0 requires no dist between chars, we just count number of chars that match the first one.
k == 1 requires the character to be one hop away. For b, either a or c is valid.
k == 2 requires the character to be two hops away. For b, either a, c, or d is valid. Cycling to z is not allowed.
k == 25 is just the length of the string.
Problem does not want the ideal string itself, just its length.

Now, if a character is not in range, what do we do? Like the f in example 1 where s = acfgbd, k = 2?
2 char hops away from a are a itself, b (1), and c(2).
2 char hops away from c are a, b, c itself, d, and e. So we skip f, no increment.
Then we use c again and g is not in range either. So we don't increment. But d is in range and we increment.
Optimal string is abcd, length 4 is the answer.

I think we can use one loop for this somehow. Just trying to decide that somehow.

We can pass forward the bounds for the char. So, for acfgbd where k = 2, we can do
     a c f g b d (k == 2)
low |a a a a a b
high|c e e e d f
cnt |1 2 2 2 3 4
*/

func Test_LongestIdealSubsequence_acfgbd_2(t *testing.T) {
	actual := longestIdealSubsequence("acfgbd", 2)

	assert.Equal(t, 4, actual)
}

func longestIdealSubsequence(s string, k int) int {
	if len(s) == 0 {
		return 0
	} else if k >= 25 {
		return len(s)
	}

	bk := byte(k)

	cnt := 1
	low, high := bounds(s[0], bk)

	fmt.Println(low, high, cnt)

	for i := 1; i < len(s); i++ {
		r := byte(s[i])

		if inRange(r, low, high) {
			low, high = bounds(r, bk)
			cnt++
		}
	}

	return cnt
}

func bounds(r byte, k byte) (byte, byte) {
	return max('a', r-k), min('z', r+k)
}

func inRange(r, low, high byte) bool {
	return r >= low && r <= high
}
