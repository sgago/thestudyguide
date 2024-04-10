package dynamic

import (
	"fmt"
	"math"
	"testing"

	"sgago/thestudyguide/col/queue"

	"github.com/stretchr/testify/assert"
)

/*
	Suppose we have a m by n matrix filled with non-negative integers,
	find a path from top left corner to bottom right corner which minimizes
	the sum of all numbers along its path.

	Note: Movements can only be either down or right at any point in time.

	Example:
		Input:
		[
			[1,3,1],
			[1,5,1],
			[4,2,1]
		]
	Output:
		7
	Explanation:
		Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
*/

func TestMinimalPathSumWithoutMemo(t *testing.T) {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	sum := minimalPathSumWithoutMemo(matrix)

	fmt.Println("sum:", sum)
	assert.Equal(t, 7, sum)
}

type stateMinSum struct {
	row   int
	col   int
	total int
}

func minimalPathSumWithoutMemo(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	totalRows := len(matrix)
	totalCols := len(matrix[0])

	if totalRows == 1 && totalCols == 1 {
		return matrix[0][0]
	}

	iter := 0
	minSum := math.MaxInt

	q := queue.New(totalRows*totalCols, stateMinSum{
		row:   0,
		col:   0,
		total: matrix[0][0],
	})

	for !q.Empty() {
		iter++
		curr := q.DeqHead()

		// Is this a solution?
		if curr.row == totalRows-1 && curr.col == totalCols-1 {
			minSum = min(minSum, curr.total)
		}

		// TODO: These ifs can be consolidated with some clever tricks
		// Move down, but only if we're not at the last row
		if curr.row < totalRows-1 {
			q.EnqTail(stateMinSum{
				row:   curr.row + 1,
				col:   curr.col,
				total: curr.total + matrix[curr.row+1][curr.col],
			})
		}

		// Move down, but only if we're not at the last column
		if curr.col < totalCols-1 {
			q.EnqTail(stateMinSum{
				row:   curr.row,
				col:   curr.col + 1,
				total: curr.total + matrix[curr.row][curr.col+1],
			})
		}
	}

	fmt.Println("iter:", iter)

	return minSum
}

func TestMinimalPathSumWithMemo(t *testing.T) {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	sum := minimalPathSumWithMemo(matrix)

	fmt.Println("sum:", sum)
	assert.Equal(t, 7, sum)
}

func minimalPathSumWithMemo(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	totalRows := len(matrix)
	totalCols := len(matrix[0])

	if totalRows == 1 && totalCols == 1 {
		return matrix[0][0]
	}

	iter := 0
	minSum := math.MaxInt

	// Init our memo
	memo := make([][]int, totalRows)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, totalCols)

		for j := 0; j < len(memo); j++ {
			memo[i][j] = math.MaxInt
		}
	}

	q := queue.New(totalRows*totalCols, stateMinSum{
		row:   0,
		col:   0,
		total: matrix[0][0],
	})

	memo[0][0] = matrix[0][0]

	for !q.Empty() {
		iter++
		curr := q.DeqHead()

		// Is this a solution?
		if curr.row == totalRows-1 && curr.col == totalCols-1 {
			minSum = min(minSum, curr.total)
		}

		// TODO: These ifs can be consolidated with some clever tricks
		/*
			So, we can help get rid of these long if blocks by storing the
			row and col deltas in arrays. Like
			deltaRow := []int{1, 0}
			deltaCol := []int{0, 1}

			Then, below we can just loop over these arrays, combine data and blocks, and make everything
			shorter. I'm leaving this as is, because, while longer, it feels easier to read
			which is all I really care about when practicing and reviewing.
		*/

		// Move down, but only if we're not at the last row
		if curr.row < totalRows-1 {
			next := stateMinSum{
				row:   curr.row + 1,
				col:   curr.col,
				total: curr.total + matrix[curr.row+1][curr.col],
			}

			memoTotal := memo[next.row][next.col]

			// So, if we got to this row and col more efficiently, we can just skip enq'ing anything.
			// If memoTotal >= next.total, then this path is more efficient.
			// If memoTotal < next.total, then this path is inefficient.
			if memoTotal >= next.total {
				memo[next.row][next.col] = next.total
				q.EnqTail(next)
			}
		}

		// Move down, but only if we're not at the last column
		if curr.col < totalCols-1 {
			next := stateMinSum{
				row:   curr.row,
				col:   curr.col + 1,
				total: curr.total + matrix[curr.row][curr.col+1],
			}

			memoTotal := memo[next.row][next.col]

			// So, if we got to this row and col more efficiently, we can just skip enq'ing anything.
			// If memoTotal >= next.total, then this path is more efficient.
			// If memoTotal < next.total, then this path is inefficient.
			if memoTotal >= next.total {
				memo[next.row][next.col] = next.total
				q.EnqTail(next)
			}
		}
	}

	fmt.Println("iter:", iter)

	return minSum
}
