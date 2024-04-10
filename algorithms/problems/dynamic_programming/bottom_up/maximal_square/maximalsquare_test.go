package dynamic

import (
	"testing"
)

/*
1, 0, 1, 0, 0
1, 0, 1, 1, 1
1, 1, 1, 1, 0
1, 0, 0, 1, 0
*/

var (
	grid = [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
	}

	memo = make(map[int]map[int]int)

	// rows = len(grid)
	// cols = len(grid[0])
)

func TestMaximalSquare(t *testing.T) {
	maximalSquare()
}

func maximalSquare() int {
	best := 0

	for col, val := range grid[0] {
		memo[0][col] = val
		best = max(best, val)
	}

	for row, val := range grid {
		memo[row][0] = val[0]
		best = max(best, val[0])
	}

	for r := range grid[1:] {
		for c := range grid[r][1:] {
			if grid[r][c] == 0 {
				continue
			}

			memo[r][c] = min(
				memo[r-1][c],
				memo[r][c-1],
				memo[r-1][c-1],
			) + 1
		}
	}

	return 0
}
