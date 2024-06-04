package triangle

import (
	"fmt"
	"math"
	"sgago/thestudyguide/col/map2d"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
The problem is to find the minimum path sum from top to bottom if given a triangle. Each step you may move to adjacent numbers on the row below.

Input
triangle: see example
Output
the minimum path sum

Examples
Example 1:
Input:
	triangle = [
		[2], // <- start here can only choose adjacent numbers
		[3,4],
		[6,5,7],
		[4,1,8,3] // <- end here with lowest score
	]
Output: 11

Explanation:
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11.

===============================================================

Mmk, so we can only pick adjacent numbers. From the ith row, we can only pick
[i, j] or [i, j+1].

Tri [i, j]
2   [0, 0]
3 4 [1, 0] [1, 1]

What are our initial conditions?
2 at tri[0][0].

What's the most basic case? It's
2

How do we get from base to 3 4? We add 2 to them.
2
[1, 0]  [1, 1]  <- coords of next
(2 + 3) (2 + 4) <- sum of next

which reduces to

2
5 6

Next level
2
5 6   <- sums from before
6 6 7 <- next row (changed from example solution to make this more exciting)

This is one with our first overlap. For subtree with 5 it's
5
[2, 0] [2, 1] <- coords of next
5 + 6  5 + 6  <- next equations
11     11     <- sum of next

for subtree with 6 it's
6
[2, 1] [2, 2] <- coords of next
6 + 6  6 + 7  <- next equations
12     10     <- sum of next

Now, we got there's overlap at [2, 1]. The solution for [2, 1] for 5 is 11; the solution for
[2, 1] for 6 is 12. We want to keep 11! It's smaller than 12. This is the overlapping subproblem.
We have two ways of getting [2, 1] via
1. [0, 0] + [1, 0] + [2, 1] = 2 + 3 + 6 = 11
2. [0, 0] + [1, 1] + [2, 1] = 2 + 4 + 6 = 12

We can keep 11 and ditch 12. Again, there's a more efficient way to get to [2, 1].

*/

var (
	tri [][]int = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
)

func TestTriangleBottomUp_2LevelTri(t *testing.T) {
	actual := bottomUp(tri[:2])
	fmt.Println(actual)
	assert.Equal(t, 5, actual)
}

func TestTriangleBottomUp_3LevelTri(t *testing.T) {
	actual := bottomUp(tri[:3])
	fmt.Println(actual)
	assert.Equal(t, 10, actual)
}

func TestTriangleBottomUp_4LevelTri(t *testing.T) {
	actual := bottomUp(tri)
	fmt.Println(actual)
	assert.Equal(t, 11, actual)
}

func bottomUp(tri [][]int) int {
	if len(tri) == 0 {
		return -1
	}

	// Mulling over a new DP structure. Seeing a lot of boilerplate with
	// looping to create the sub arrays/maps, setting init values if any,
	// checking if map values exists before using them.
	dp := map2d.NewDp[int, int](2)

	dp.Set(0, 0, tri[0][0])

	for i, lvl := range tri[:len(tri)-1] {
		i = i + 1

		for j := range lvl {
			prev := dp[i-1][j]
			leftNum := tri[i][j] + prev
			rightNum := tri[i][j+1] + prev

			dp.
				Min(i, j, leftNum).
				Min(i, j+1, rightNum)
		}

		// Lesson learned: Watch out, you don't need to keep
		// the entire dp structure. Here, we just need to keep the last
		// dp numbers we crunched.
		dp.Del(i - 1)
	}

	best := math.MaxInt
	last, _ := dp.SubMap(len(tri) - 1)

	for _, val := range last {
		best = min(best, val)
	}

	return best
}
