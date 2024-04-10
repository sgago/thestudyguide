package longest_common_subsequence

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"
)

/*
	Longest common subsequence (LCS)
	https://algo.monster/problems/longest_common_subsequence

	Our memo could be a 2D array. We usually memo the same
	thing as the output, so we'd memo the LCS for a given letter.

	memo[i][j] = holds max LCS at i,j.
	memo[i][j+1] = when char[i]==char[j+1], then
		memo[i][j+1]=max(memo[i][j+1], curr+1)

	  a a b c  a e
    a 1 2 0 0  1 0
	c 0 0 0 3? 0 0
	e 0 0 0 0  0 4?

	aace?
*/

func TestLcsWithMemo(t *testing.T) {
	word1, word2 := "aabcae", "ace"

	actual := lcsBfsWithMemo(word1, word2)
	fmt.Println("lcs:", actual)
}

type state struct {
	i   int
	j   int
	len int
}

func lcsBfsWithMemo(w1, w2 string) int {
	q := queue.New[state](100)

	memo := make([][]int, len(w1))
	for i := 0; i < len(w1); i++ {
		memo[i] = make([]int, len(w2))
	}

	q.EnqTail(state{
		i:   -1,
		j:   0,
		len: -1,
	})

	soln := 0
	iter := 0

	for !q.Empty() {
		curr := q.DeqHead()

		if curr.i == len(w1)-1 {
			continue
		}

		nextI := curr.i + 1
		if curr.i >= len(w1)-1 {
			nextI = 0
		}

		for j := 0; j < len(w2); j++ {
			iter++

			nextLen := curr.len
			if w1[nextI] == w2[j] {
				nextLen++
			}

			soln = max(soln, nextLen)

			// Have we crunched nextI,j before? Did we get to nextI,j in a more efficient way?
			// If so, let's not enqueue anything else cause there's no point.
			if cnt := memo[nextI][j]; cnt > nextLen {
				continue
			}

			memo[nextI][j] = nextLen

			q.EnqTail(state{
				i:   nextI,
				j:   j,
				len: nextLen,
			})
		}
	}

	fmt.Println("iter:", iter)
	return soln
}
