package wordbreakii

import (
	"fmt"
	"sgago/thestudyguide/col/queue"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
140. Word Break II
https://leetcode.com/problems/word-break-ii

Given a string s and a dictionary of strings wordDict,
add spaces in s to construct a sentence where each word is a
valid dictionary word. Return all such possible sentences in any order.

Note that the same word in the dictionary may be reused
multiple times in the segmentation.

Example 1:
	Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
	Output: ["cats and dog","cat sand dog"]

Example 2:
Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: []

Constraints:
- 1 <= s.length <= 20
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 10
- s and wordDict[i] consist of only lowercase English letters.
- All the strings of wordDict are unique.
- Input is generated in a way that the length of the answer doesn't exceed 105.
*/

/*
Mmk, add spaces to the string s using the words in the wordDict.
Given "catdog" and ["cat", "dog"] return ["cat dog"].

DFS out solutions.
*/

func TestWordBreak2Brute(t *testing.T) {
	actual := wordBreak2Brute("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})

	fmt.Println(actual)

	assert.Len(t, actual, 2)
	// assert.Contains(t, actual, "cats and dog")
	// assert.Contains(t, actual, "cat sand dog")
}

func wordBreak2Brute(s string, wordDict []string) [][]string {
	soln := make([][]string, 0)
	spaces := make([][]string, 0)

	for _, word := range wordDict {
		if ok := strings.HasPrefix(s, word); ok {
			spaces = append(spaces, []string{word})
		}
	}

	n := 0

	for {
		cnt := len(spaces[0])
		last := spaces[0][len(spaces[0])-1]
		n = len(last) + n

		for _, word := range wordDict {

			if ok := strings.HasPrefix(s[n:], word); ok {
				spaces[0] = append(spaces[0], word)

				if n+len(word) == len(s) {
					soln = append(soln, spaces[0])
					spaces = slices.Delete(spaces, 0, 1)
					cnt = -1
					n = 0
				}

				break
			}
		}

		if len(spaces) == 0 {
			break
		}

		if cnt == len(spaces[0]) {
			spaces = slices.Delete(spaces, 0, 1)
		}
	}

	return soln
}

func TestWordBreak2Dfs(t *testing.T) {
	actual := wordBreak2Dfs("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})

	fmt.Println(actual)

	assert.Len(t, actual, 2)
	assert.Contains(t, actual, "cats and dog")
	assert.Contains(t, actual, "cat sand dog")
}

type state struct {
	s      string
	result string
}

func wordBreak2Dfs(s string, wordDict []string) []string {
	ans := make([]string, 0)

	if s == "" {
		return ans
	} else if len(wordDict) == 0 {
		return ans
	}

	q := queue.New[state](len(wordDict))

	for _, word := range wordDict {
		// TODO: I wonder if stringbuilder would be faster?
		// CutPrefix is convenient for this.
		if after, ok := strings.CutPrefix(s, word); ok {
			q.EnqTail(state{
				s:      after,
				result: word,
			})
		}
	}

	for !q.Empty() {
		curr := q.DeqHead()

		if curr.s == "" {
			ans = append(ans, curr.result)
			continue
		}

		for _, word := range wordDict {
			if after, ok := strings.CutPrefix(curr.s, word); ok {
				q.EnqTail(state{
					s:      after,
					result: curr.result + " " + word,
				})
			}
		}
	}

	return ans
}
