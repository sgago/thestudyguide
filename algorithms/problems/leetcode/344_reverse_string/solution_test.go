package reversestring

import (
	"fmt"
	"testing"
)

/*
344. Reverse String
https://leetcode.com/problems/reverse-string

Write a function that reverses a string.
The input string is given as an array of characters s.

You must do this by modifying the input array
in-place with O(1) extra memory.



Example 1:
	Input: s = ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]

Example 2:
	Input: s = ["H","a","n","n","a","h"]
	Output: ["h","a","n","n","a","H"]


Constraints:
- 1 <= s.length <= 105
- s[i] is a printable ascii character.
*/

func TestReverseString_Hello(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)

	fmt.Println(s)
}

func reverseString(s []byte) {
	// TODO Check params

	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]

		i++
		j--
	}
}
