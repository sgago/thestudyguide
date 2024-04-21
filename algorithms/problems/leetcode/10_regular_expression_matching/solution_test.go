package regularexpressionmatching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
10. Regular Expression Matching
https://leetcode.com/problems/regular-expression-matching/

Given an input string s and a pattern p, implement regular
expression matching with support for '.' and '*' where:
- '.' Matches any single character.​​​​
- '*' Matches zero or more of the preceding element.
- The matching should cover the entire input string (not partial).

Example 1:
	Input: s = "aa", p = "a"
	Output: false
	Explanation: "a" does not match the entire string "aa".

Example 2:
	Input: s = "aa", p = "a*"
	Output: true
	Explanation: '*' means zero or more of the preceding element, 'a'.
		Therefore, by repeating 'a' once, it becomes "aa".

Example 3:
	Input: s = "ab", p = ".*"
	Output: true
	Explanation: ".*" means "zero or more (*) of any character (.)".

Constraints:
- 1 <= s.length <= 20
- 1 <= p.length <= 20
- s contains only lowercase English letters.
- p contains only lowercase English letters, '.', and '*'.
- It is guaranteed for each appearance of the character '*',
  there will be a previous valid character to match.

===============================================

Mmk, so, there's a very low length to strings s and p, 20 each. Leetcode
never does this out of the "goodness of the hearts", so it *could*
suggest DP-like TC.

What are our options here?
- We can do a two pointer system, a slow + fast when looking forward to any *'s that appear.
- I don't envision any monotonicity to the inputs. They are effectively random.
- Weird inputs like .*.*.* is effectively .*. Maybe the "previous valid character match" constraint stops this.
- Reminder, * is zero to many.
- We can basically mountain climb this. * is easily the hardest part, so let's start with that.
- After building *, we can add in dots, then add in .* checks.
- Would DFS help like at all? What would we be searching for anyway? This doesn't feel dfs-y yet.
- Would a stack like structure help somehow? How would this be different than two pointers?
- Note that only a matching character to * will stop the lookahead. Like a* will only be stopped by a b or something.
  Likewise, .* is only stopped by a character or single.
- All I see is a two-pointer forward checks, walking along the string checking for *'s, and chewing thru the problem that way.
- Hmm, what about string split? Would that maybe help? IDK.

for chars in str and pattern:
  Do the * lookahead
  Is the next char end of string or a *?
	If not, do a "simple" char check. Adv slow and fast ptrs.
  	Else determine what would stop our * consumption,
	  consume chars until you hit that. a* is stopped by new chars like b, c, or end of str.

Pause here. Am I on the right path with two pointer + lookahead * work?

Aww man, DP was the winning solution. Completely called it earlier due to the length of the strings.
Should have listened to my intuition lol. The other give away here is
two inputs or sequences, that'll yield a 2D problem.

Mmk, so, this is DP then. Welp, let's go then. Memo will hold true/false. Note, can be bit storage then.
Start state is true.
Input is a and we've got a*. Init state is true with nothing.

abcaaa  abca*



*/

func TestIsMatch_aa_a_ReturnsFalse(t *testing.T) {
	actual := isMatch("a", "aa")

	assert.False(t, actual)
}

func TestIsMatch_aa_astar_ReturnsFalse(t *testing.T) {
	actual := isMatch("a", "a*")

	assert.True(t, actual)
}

func isMatch(str, pattern string) bool {

	return false
}
