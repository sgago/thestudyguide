package removealloccurencesofasubstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1910. Remove All Occurrences of a Substring
https://leetcode.com/problems/remove-all-occurrences-of-a-substring/

Given two strings s and part, perform the following operation on s until
all occurrences of the substring part are removed:
- Find the leftmost occurrence of the substring part and remove it from s.
Return s after removing all occurrences of part.

A substring is a contiguous sequence of characters in a string.

Example 1:
	Input: s = "daabcbaabcbc", part = "abc"
	Output: "dab"
	Explanation: The following operations are done:
	- s = "daabcbaabcbc", remove "abc" starting at index 2, so s = "dabaabcbc".
	- s = "dabaabcbc", remove "abc" starting at index 4, so s = "dababc".
	- s = "dababc", remove "abc" starting at index 3, so s = "dab".
	Now s has no occurrences of "abc".

Example 2:
	Input: s = "axxxxyyyyb", part = "xy"
	Output: "ab"
	Explanation: The following operations are done:
	- s = "axxxxyyyyb", remove "xy" starting at index 4 so s = "axxxyyyb".
	- s = "axxxyyyb", remove "xy" starting at index 3 so s = "axxyyb".
	- s = "axxyyb", remove "xy" starting at index 2 so s = "axyb".
	- s = "axyb", remove "xy" starting at index 1 so s = "ab".
	Now s has no occurrences of "xy".

Constraints:
- 1 <= s.length <= 1000
- 1 <= part.length <= 1000
- s​​​​​​ and part consists of lowercase English letters.

=======================================================================

There's many ways to solve this one, on first glance.
- We can use strings.ReplaceAll function, of course.
- The brute approach will compare letters and remove the part.
- An improvement would be to compare the last letter of part and work backwards.
  - If there's no match, you can keep moving backwards to find the start of part to save time.
  - If there's a match, you can remove the part and continue.
- We can use a stack to keep track of the letters and compare the last letter of part.
- We can use a two-pointer approach to compare the letters and remove the part.
- We can use a sliding window approach to compare the letters and remove the part.
- We can use a regex to compare the letters and remove the part.
- Rabin-Karp or KMP algorithm can be used to compare the letters and remove the part.

A key part here, is that we need to remove all occurrences of the substring part.
Removing a part can create a new part!
For example, removing "abc" from aabcbc will create a new "abc" that we need to remove.
*/

func TestRemoveAllOccurrencesOfASubstringReplaceAll(t *testing.T) {
	actual := RemoveAllOccurrencesOfASubstringReplaceAll("daabcbaabcbc", "abc")
	assert.Equal(t, "dab", actual)
}

func RemoveAllOccurrencesOfASubstringReplaceAll(s, part string) string {

	for strings.Contains(s, part) {
		s = strings.ReplaceAll(s, part, "")
	}

	return s
}
