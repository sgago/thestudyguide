package lettertilepossibilities

import (
	"sgago/thestudyguide/col/flags"
	"sgago/thestudyguide/col/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1079. Letter Tile Possibilities
https://leetcode.com/problems/letter-tile-possibilities/

You have n tiles, where each tile has one letter tiles[i] printed on it.

Return the number of possible non-empty sequences of letters you
can make using the letters printed on those tiles.

Example 1:
	Input: tiles = "AAB"
	Output: 8
	Explanation: The possible sequences are
		"A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
	Example 2:

Input: tiles = "AAABBC"
	Output: 188
	Example 3:

Input: tiles = "V"
	Output: 1


Constraints:
- 1 <= tiles.length <= 7
- tiles consists of uppercase English letters.


================================== Solution Approach ==================================

Mmk, we have to create all unique sequences of letters from the given tiles.
For example, "AAB" can't create "A" and "A" again. We can use
DFS, backtracking, and memoization to solve this one.
We'll need a solution map to store seen solutions and a used map to avoid
repeating letters.

Can we use dynamic programming to solve this problem? I think we can do that, too.
This is also a permutations problem. We can math it out to solve this one quickly.
This is called combinatorics. We can use the formula like n! / (n1! * n2! * ... * nk!)
where n is the total number of tiles, and n1, n2, ..., nk are the number of characters.
*/

func TestLetterTilePossibilitiesDfs(t *testing.T) {
	tiles := "AAB"
	actual := LetterTilePossibilitiesDfs(tiles)
	assert.Equal(t, 8, actual)
}

type state struct {
	// tile is the current string of tiles (letters).
	tile string

	// used keeps track of which letters have been used.
	used flags.Flags
}

func LetterTilePossibilitiesDfs(tiles string) int {
	if len(tiles) <= 1 {
		return len(tiles)
	}

	soln := make(map[string]bool)

	init := state{
		tile: "",
		used: flags.New(len(tiles)),
	}

	stack := stack.New(10, init)

	for !stack.Empty() {
		curr := stack.Pop()

		if _, ok := soln[curr.tile]; ok {
			continue
		}

		soln[curr.tile] = true

		for i, r := range tiles {
			if curr.used.Get(i) {
				continue
			}

			next := state{
				tile: curr.tile + string(r),
				used: curr.used.Clone(),
			}

			next.used.True(i)

			stack.Push(next)
		}
	}

	return len(soln) - 1
}
