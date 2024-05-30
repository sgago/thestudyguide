package findallanagramsinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
438. Find All Anagrams in a String
https://leetcode.com/problems/find-all-anagrams-in-a-string

Given two strings s and p, return an array of all the start
indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters
of a different word or phrase, typically using all the original
letters exactly once.


Example 1:
	Input: s = "cbaebabacd", p = "abc"
	Output: [0,6]
	Explanation:
		The substring with start index = 0 is "cba", which is an anagram of "abc".
		The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
	Input: s = "abab", p = "ab"
	Output: [0,1,2]
	Explanation:
		The substring with start index = 0 is "ab", which is an anagram of "ab".
		The substring with start index = 1 is "ba", which is an anagram of "ab".
		The substring with start index = 2 is "ab", which is an anagram of "ab".

Constraints:
- 1 <= s.length, p.length <= 3 * 104
- s and p consist of lowercase English letters.
*/

/*
So, we want to find the starting index of all anagrams in s by rearranging
all the chars in p and finding a substring that matches.

Brute is a lot of looping, rearranging and retrying.

Is this just a two pointer? We sum up the bytes in p and find a matching
sum over in s? abc sums to 6 and so does cba. bae sums to like 10 and 6 != 10.
aaa = 3, aba = 4. Yeah, sum p up and do prefix sum like soln over in s.
Gets us m+n TC.
*/

func TestFindAllAnagramsInAString_cbaebabacd_abc(t *testing.T) {
	actual := findAllAnagramsInAString("cbaebabacd", "abc")
	fmt.Println(actual)

	assert.Len(t, actual, 2)
	assert.Contains(t, actual, 0)
	assert.Contains(t, actual, 6)
}

func TestFindAllAnagramsInAString_cbaebabacd_a(t *testing.T) {
	actual := findAllAnagramsInAString("cbaebabacd", "a")
	fmt.Println(actual)

	assert.Len(t, actual, 3)
	assert.Contains(t, actual, 2)
	assert.Contains(t, actual, 5)
	assert.Contains(t, actual, 7)
}

func findAllAnagramsInAString(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	checksum := 0

	for _, r := range p {
		checksum += rtoi(r)
	}

	ans := []int{}

	sum := 0   // Window sum
	start := 0 // 2 ptr start
	end := 0   // 2 ptr end

	for ; end < len(p); end++ {
		sum += btoi(s[end])
	}

	end-- // Loop stops with end at 3, we need it at 2

	for end < len(s)-1 {
		if sum == checksum {
			ans = append(ans, start)
		}

		sum -= btoi(s[start])
		start++

		end++
		sum += btoi(s[end])
	}

	return ans
}

func rtoi(r rune) int {
	return int(r - ('a' - 1))
}

func btoi(b byte) int {
	return int(b - ('a' - 1))
}
