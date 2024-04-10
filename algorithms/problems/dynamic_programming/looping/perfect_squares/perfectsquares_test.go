package dynamic

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"

	"github.com/stretchr/testify/assert"
)

/*
Perfect Squares
https://algo.monster/problems/perfect_squares

Given a number that is less than 10^5 determine the smallest
amount of perfect squares needed to sum to a particular number.
The same number can be used multiple times.

A perfect square is an integer p that can be represented as q^2
for some other integer q. For example, 9 is a perfect square since 9 = 3^2.
However, 8 isn't a perfect square.

Examples

Example 1:
	Input: 12
	Output: 3
Explanation:
	12 = 4 + 4 + 4

Example 2:
	Input: 13
	Output: 2
Explanation:
	13 = 4 + 9

Mmk, so
1. Perfect squares (psqrs) are anything that's X^2: 4, 9, 16, 25, etc.
2. We want to sum up as few perfect squares as possible.
3. We want to sum these squares to hit some target.
4. So, given 16, 4*4 = 4^2 = 16. Done.
5. Given 15, 3*3 = 3^2 = 9. 2*2 = 2^2 = 4. At 13...,
   1 and 0 are also perfect squares, we can just increment by 2 = 1^2 + 1^2 = 1 + 1.
6. I don't just "know" what the perfect square values are.
7. We can do log magic to try to help find answers too.
8. We can also do stuff like 144^(1/2) = 12. There's no guarantee of a min number of psqrs while doing this.

So, I guess we can generate all psqrs <= target. For 12 that would be 1, 4, 9.
For diffs of 3, we instantly add 3 to the our total psqrs used.

target = psqr(a) + psqr(b) + ... psqr(n)
target^(0.5) = (psqr(a) + psqr(b) + ... psqr(n))^(0.5)

Remember that we can reuse numbers, hence the soln for 12 being 4+4+4.

If I were to DFS/BFS this monster for 12, it'd be like

All psqrs to 12 = 1, 4, 9. These are the nums we get to work with. Funny, we could also gen, the doubled pqrs
all psqrs*2 <= target which would be 1, 8 and also
all psqrs*3 <= target which would be 1, 12
all psgrs*4 <= target which would be just 1, so stop here...

Reminder output is just a count.

12
-1 -4 -9 (subtract from above, | are for readability)
11  8  3 (result from subtraction, pqrs count is 1 to this level)
-1 -4 -9 | -1 -4 -9 | -1 -4 -9 (subtract from above)
10  7  2 |  7  4 -6 |  2 -1 -6 (result from subtraction, pqrs count is 2 at this level)
10  7  2 |  7  4 |     2       (remove negatives, they are not solns)
-1 -4 -9 | -1 -4 -9 | -1 -4 -9 | -1 -4 -9 | -1 -4 -9 | -1 -4 -9 (subtract from above)
....... | ...... | 0! ....... (pqrs count is 3 at this level)

Same tree without the subtraction noise:
12
11 | 8 | 3 (result from subtraction, pqrs count is 1 to this level)
10 7 2 | 7 4 | 2 (remove negatives, they are not solns)
....... | ...... | 0 ....... (pqrs count is 3 at this level)

We want to BFS, not DFS. The soln is helped by broad search, not deep search.
This has a major 2^N feel about it. Can we memoize here? What would we memoize?

We could just memoize any non-solns found and remove those from our state space, I guess.
In above, only 7 appears twice. We don't end up doing this but we could say 7=false, not a soln, and skip it.
Or, we've already got one 7, so we could just trim it off too with the negatives. Maybe not, could get bad count...

For DP problems, we basically always seem to be memoizing the solns to the subproblems,
so the same thing as whatever the final solution represents. If it's the length of the longest subsequence,
then we store the length of each subproblem. Same with largest div subset, it's also length.

So, here, our output is a count. So we could just memo a count. This is the thing I need
to remember, 100% with these problems. The memo typically stores the same thing as the output.
So, what we could do, is memo the count of sqrs it takes to reach a particular number.
In this case,.... oooo we could do memo[i] = min(moves to get to any i).
Then memo[i+1] = min(memo[i]).

min moves to 1 is 1, 4 is 1, 8 is 1, etc.
*/

func TestPerfectSquares_15(t *testing.T) {
	actual := perfectSquares(15)
	fmt.Println(" cnt:", actual)
	assert.Equal(t, 4, actual)
}

type state struct {
	count int
	total int
	str   string
}

func perfectSquares(target int) int {
	if target == 0 {
		return 1
	}

	sqrs := []int{}
	q := queue.New[state](10)
	memo := map[int]int{}

	// Unlike other DP problems, we need to build
	// our array. For this problem we build an array of
	// perfect squares (sqr and sqrs).
	for i := 0; i*i <= target; i++ {
		sqr := i * i

		if sqr == target {
			return 1
		}

		memo[sqr] = 1
		sqrs = append(sqrs, sqr)

		q.EnqHead(state{
			count: 1,
			total: sqr,
			str:   fmt.Sprint(sqr),
		})
	}

	iter := 0

	for !q.Empty() {
		curr := q.DeqHead()

		for _, sqr := range sqrs {
			iter++
			next := curr.total + sqr
			count := curr.count + 1

			// Is next a solution?
			if next == target {
				fmt.Println("iter:", iter)
				fmt.Println(" equ:", curr.str+"+"+fmt.Sprint(sqr))
				return count
			}

			// Is next > target? If so, there's no solution
			// so continue on to the next value in our queue.
			if next > target {
				continue
			}

			// Have we looked at this next value before?
			// Did we get to the next value with a lower count?
			// If so, there's no point in enq'ing the next round of
			// perfect sqrs. We got to next in some more efficient
			// way, so we can just continue on to the next number.
			if memoCount, ok := memo[next]; ok && count >= memoCount {
				continue
			}

			memo[next] = curr.count

			q.EnqTail(state{
				count: curr.count + 1,
				total: curr.total + sqr,
				str:   curr.str + "+" + fmt.Sprint(sqr),
			})
		}
	}

	return -1
}
