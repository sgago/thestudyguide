/*
Partitioning A String Into Palindromes
https://algo.monster/problems/palindrome_partitioning

Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Examples
Example 1:
Input: aab
Output:
	a|a|b
	aa|b

Examples
Example 2:
Input: aabob
Output:
	a|a|b|o|b
	aa|b|o|b
	a|a|bob
	aa|bob
*/

package backtracking

import (
	"testing"

	"sgago/thestudyguide/col/node"
	stringutil "sgago/thestudyguide/utils/strings"
)

func TestPalindromePartitions(t *testing.T) {
	str := "aab"

	n := node.New[string](str)

	for i := len(str); i > 0; i-- {
		if stringutil.Palindrome(str[:i]) {
			n.AddChild(str[:i])
		}
	}
}

func pal(n *node.Node[string]) {
	if len(n.Val) == 0 {
		return
	}

	for i := len(n.Val); i > 0; i-- {
		if stringutil.Palindrome(n.Val[:i]) {
			n.AddChild(n.Val[:i])
		}
	}

}
