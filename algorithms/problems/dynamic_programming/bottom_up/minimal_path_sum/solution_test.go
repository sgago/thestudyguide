package minimalpathsum

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Minimal Path Sum
	https://algo.monster/problems/minimal_path_sum

	Lesson learned: Think I might be starting to see the light. Bottom-up DP problems
	are calculating the next values from previous one(s). Here, we look at the previous
	values up and left, pick the min value, and current value, and move onto the next.
*/

func TestMinimalPathSum(t *testing.T) {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	// Min is 7 = 1 -> 3 -> 1 -> 1 -> 1
	actual := minimalPathSum(matrix)

	assert.Equal(t, 7, actual)
}

// TC is r*c
func minimalPathSum(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	// Memo will store the cheapest path to the bottom right
	dp := make([][]int, rows)
	for i := 0; i < cols; i++ {
		dp[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		// Loop moves us down

		for c := 0; c < cols; c++ {
			// Loop moves us right
			curr := matrix[r][c]
			dp[r][c] = math.MaxInt

			if r == 0 && c == 0 {
				dp[r][c] = curr
				continue
			}

			up := math.MaxInt
			left := math.MaxInt

			if r > 0 {
				up = dp[r-1][c]
			}

			if c > 0 {
				left = dp[r][c-1]
			}

			dp[r][c] = min(up, left) + curr
		}
	}

	return dp[rows-1][cols-1]
}
