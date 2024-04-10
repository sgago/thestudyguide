package gcdofstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	1071. Greatest Common Divisor of Strings
	https://leetcode.com/problems/greatest-common-divisor-of-strings/

	For two strings s and t, we say "t divides s" if and only if
	s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

	Given two strings str1 and str2, return the largest
	string x such that x divides both str1 and str2.

	Example 1:
		Input: str1 = "ABCABC", str2 = "ABC"
		Output: "ABC"

	Example 2:
		Input: str1 = "ABABAB", str2 = "ABAB"
		Output: "AB"

	Example 3:
		Input: str1 = "LEET", str2 = "CODE"
		Output: ""

	Constraints:
		1 <= str1.length, str2.length <= 1000
		str1 and str2 consist of English uppercase letters.

	===============================================================================

	Mmk. So, the GCD in math is like the greatest value two numbers have in common
	that can be divided by a whole positive integer. GCD of 54 and 24 is 6.

	With that in mind, we're trying to find the longest string that fits into
	both str1 and str2. One observation is that we can use the string lengths.
	Say lengths of str1 and str2 are n1 and n2 respectively. For example,
	n1 == 3 and n2 == 2. The biggest gcd we can get is 1. For Input: str1 = "ABABAB", str2 = "ABAB",
	n1 == 6 and n2 == 4. There's no point checking ABAB because n2 * 2 = 8 which is > n1.
	There's also no point in checking ABA because 3 * 2 = 6 > n2. Both strings need modulo == 0 to work out
	with our proposed result.

	If the initial characters don't match, we can also quit early.
*/

func TestGcdOfStrings_ABCABC_ABC(t *testing.T) {
	actual := gcdOfStrings("ABCABC", "ABC")
	assert.Equal(t, "ABC", actual)
}

func TestGcdOfStrings_ZBC_ABC(t *testing.T) {
	actual := gcdOfStrings("ZBCABC", "ABC")
	assert.Equal(t, "", actual)
}

func TestGcdOfStrings_ABABAB_ABAB(t *testing.T) {
	actual := gcdOfStrings("ABABAB", "ABAB")
	assert.Equal(t, "AB", actual)
}

func TestGcdOfStrings_LEET_CODE(t *testing.T) {
	actual := gcdOfStrings("LEET", "CODE")
	assert.Equal(t, "", actual)
}

func gcdOfStrings(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)

	if n1 == 0 || n2 == 0 {
		// If either string is empty, quit early.
		return ""
	}

	if str1[0] != str2[0] {
		// Quick shortcut. If first letters are not equal, then
		// we can just quit early, too. Neither is divisible by the other.
		return ""
	}

	for size := min(n1, n2); size > 0; size-- {
		if n1%size > 0 || n2%size > 0 {
			// Need both str1 and str2 to be divisible
			// by our current size. Otherwise, there's no point in checking.
			continue
		}

		maybe := str2[:size] // Maybe a solution, just pick str2 but could be either, I think.

		if divisible(str1, maybe) && divisible(str2, maybe) {
			return maybe
		}
	}

	return ""
}

func divisible(quotient, divisor string) bool {
	n := len(divisor)

	for i := 0; i < len(quotient); i += n {
		if quotient[i:i+n] != divisor {
			return false
		}
	}

	return true
}
