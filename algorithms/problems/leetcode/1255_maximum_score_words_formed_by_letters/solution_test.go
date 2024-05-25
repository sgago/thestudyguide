package maximumscorewordsformedbyletters

import (
	"fmt"
	"sgago/thestudyguide/col/queue"
	"testing"
)

/*
Maximum Score Words Formed By Letters
https://leetcode.com/problems/maximum-score-words-formed-by-letters

Given a list of words, list of  single letters (might be repeating), and
score of every character.

Return the maximum score of any valid set of words formed by using the
given letters (words[i] cannot be used two or more times).

It is not necessary to use all characters in letters and each letter can
only be used once. Score of letters 'a', 'b', 'c', ... ,'z' is given by score[0],
score[1], ... , score[25] respectively.

Example 1:
	Input:
		words = ["dog","cat","dad","good"],
		letters = ["a","a","c","d","d","d","g","o","o"],
		score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
	Output: 23
	Explanation:
		Score  a=1, c=9, d=5, g=3, o=2
		Given letters, we can form the words "dad" (5+1+5) and
		"good" (3+2+2+5) with a score of 23.
		Words "dad" and "dog" only get a score of 21.

Example 2:
	Input:
		words = ["xxxz","ax","bx","cx"],
		letters = ["z","a","b","c","x","x","x"],
		score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
	Output: 27
	Explanation:
		Score  a=4, b=4, c=4, x=5, z=10
		Given letters, we can form the words "ax" (4+5), "bx" (4+5) and
		"cx" (4+5) with a score of 27.
		Word "xxxz" only get a score of 25.

Example 3:
	Input:
		words = ["leetcode"],
		letters = ["l","e","t","c","o","d"],
		score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
	Output: 0
	Explanation:
		Letter "e" can only be used once.

Constraints:
- 1 <= words.length <= 14
- 1 <= words[i].length <= 15
- 1 <= letters.length <= 100
- letters[i].length == 1
- score.length == 26
- 0 <= score[i] <= 10
- words[i], letters[i] contains only lower case English letters.
*/

/*
Mmk, so, we need to use the letters in letters to form the words
in words that gives us the highest score per letter used.
We can only use each word once. We must be able to form the
entire word or words using the letters provided.

Words and letters are not guaranteed to be sorted.
Words and letters lengths are pretty short, suggesting
DFS or DP.

We could enter letters, letter counts, and scores into
a dictionary. We could then determine if we can build
each word and it's score.

This is probably DP. We're trying to maximize score
and we've got overlapping sub problems and we can only
use letters a maximum number of times.

What does our memo look like?

Score  a=1, c=9, d=5, g=3, o=2
words = ["dog","cat","dad","good"],
letters = ["a","a","c","d","d","d","g","o","o"],

           key=cnt,score
scoreMap = a=2,1 c=1,9 d=3,5 g=1,3 0=2,2

  dog  cat  dad  good
  523  92x  515  3225
  10   0    11   12

    d o g | c a t | d a d | g o o d
2 a          x21 0
1 c         19
3 d 5
1 g     10
2 o   7
*/

func TestMaximumScore(t *testing.T) {
	words := []string{"dog", "cat", "dad", "good"}
	letters := []string{"a", "a", "c", "d", "d", "d", "g", "o", "o"}
	scores := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	score := maxScore(words, letters, scores)

	fmt.Println(score)
}

type lsc struct {
	score int
	count int
}

type state struct {
	used     []bool // OPTI: Bitmask this
	score    int
	scoreMap map[string]lsc
}

func maxScore(words []string, letters []string, scores []int) int {
	// TODO: Check params

	ans := 0
	scoreMap := make(map[string]lsc, 26)

	// TC O(L)
	for _, letter := range letters {
		r := letter[0] - 'a'

		if sm, ok := scoreMap[letter]; !ok {
			scoreMap[letter] = lsc{
				score: scores[int(r)],
				count: 1,
			}
		} else {
			sm.count++
		}
	}

	fmt.Println(scoreMap)

	// We need to concat each word, determining if we can
	// use the letters, and DFS out the max score.

	q := queue.New[state](4)

	for i, word := range words {
		// OPTI: Remove words we can't make

		used := make([]bool, len(words))
		used[i] = true

		score, ok, sm := scoreWord(word, scoreMap)

		if ok {
			q.EnqTail(state{
				used:     used,
				score:    score,
				scoreMap: sm,
			})
		}
	}

	for !q.Empty() {
		curr := q.DeqHead()
		ans = max(curr.score, ans)

		for i := 0; i < len(curr.used); i++ {
			if curr.used[i] {
				continue
			}

			score, ok, sm := scoreWord(words[i], curr.scoreMap)

			if !ok {
				continue
			}

			next := state{
				used:     curr.used,
				score:    score + curr.score,
				scoreMap: sm,
			}

			next.used[i] = true

			q.EnqTail(next)
		}
	}

	return ans
}

func scoreWord(word string, scoreMap map[string]lsc) (int, bool, map[string]lsc) {
	ans := 0

	for _, l := range word {
		letter := string(l)

		sm, ok := scoreMap[letter]

		if !ok {
			// Letter does not exist in the score map
			// We cannot create the word
			// Exit now, returning false
			return -1, false, scoreMap
		} else if sm.count == 0 {
			// We ran out of letters
			// We cannot make this word with the provided score map
			// Exit now, returning false
			return -1, false, scoreMap
		}

		ans += sm.score
		sm.count--
	}

	return ans, true, scoreMap
}
