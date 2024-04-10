package mergestringsalternatively1768

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1768. Merge Strings Alternately
https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

1/75 of the leetcode 75.

You are given two strings word1 and word2.
Merge the strings by adding letters in alternating order,
starting with word1. If a string is longer than the other,
append the additional letters onto the end of the merged
string.

Return the merged string.

Example 1:
	Input: word1 = "abc", word2 = "pqr"
	Output: "apbqcr"
	Explanation: The merged string will be merged as so:
	word1:  a   b   c
	word2:    p   q   r
	merged: a p b q c r

Example 2:
	Input: word1 = "ab", word2 = "pqrs"
	Output: "apbqrs"
	Explanation: Notice that as word2 is longer, "rs" is appended to the end.
	word1:  a   b
	word2:    p   q   r   s
	merged: a p b q   r   s

Example 3:
	Input: word1 = "abcd", word2 = "pq"
	Output: "apbqcd"
	Explanation: Notice that as word1 is longer, "cd" is appended to the end.
	word1:  a   b   c   d
	word2:    p   q
	merged: a p b q c   d


Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.
*/

func TestMergeStringsAlternatively_SameLength(t *testing.T) {
	word1, word2 := "abc", "pqr"
	actual := mergeStringsAlternatively(word1, word2)
	assert.Equal(t, "apbqcr", actual)
}

func TestMergeStringsAlternatively_Word1LongerThanWord2(t *testing.T) {
	word1, word2 := "abcdef", "pqr"
	actual := mergeStringsAlternatively(word1, word2)
	assert.Equal(t, "apbqcrdef", actual)
}

func TestMergeStringsAlternatively_Word2LongerThanWord1(t *testing.T) {
	word1, word2 := "abc", "pqrstuv"
	actual := mergeStringsAlternatively(word1, word2)
	assert.Equal(t, "apbqcrstuv", actual)
}

func TestMergeStringsAlternatively_Word1IsEmpty(t *testing.T) {
	word1, word2 := "", "pqrstuv"
	actual := mergeStringsAlternatively(word1, word2)
	assert.Equal(t, word2, actual)
}

func TestMergeStringsAlternatively_Word2IsEmpty(t *testing.T) {
	word1, word2 := "abcdef", ""
	actual := mergeStringsAlternatively(word1, word2)
	assert.Equal(t, word1, actual)
}

func mergeStringsAlternatively(word1, word2 string) string {
	minLen := min(len(word1), len(word2))
	mix := make([]rune, minLen*2)

	i := 0

	for ; i < minLen; i++ {
		mix[i*2] = rune(word1[i])
		mix[i*2+1] = rune(word2[i])
	}

	return string(mix) + word1[i:] + word2[i:]
}
