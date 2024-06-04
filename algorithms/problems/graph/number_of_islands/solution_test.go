package graph

import (
	"testing"

	"sgago/thestudyguide/col/graph"
	"sgago/thestudyguide/col/queue"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfIslands_2x2(t *testing.T) {
	matrix := [][]int{
		{0, 1},
		{1, 0},
	}

	islands := numberOfIslands(matrix)
	assert.Equal(t, 2, islands)
}

func TestNumberOfIslands_5x6(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}

	islands := numberOfIslands(matrix)
	assert.Equal(t, 2, islands)
}

func TestNumberOfIslands_3x3_TotalIsland(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	islands := numberOfIslands(matrix)
	assert.Equal(t, 1, islands)
}

func TestNumberOfIslands_3x3_NoIsland(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	islands := numberOfIslands(matrix)
	assert.Equal(t, 0, islands)
}

type state struct {
	row int
	col int
}

// numberOfIslands counts the number of islands in a matrix, that is,
// groups of ones surrounded by zeros.
// Diagonals don't count, only up, down, left, and right.
//  1. Marks all zero's as visited immediately.
//  2. Ignores already visited elements.
//  3. If a one is encountered, increments count, then runs a BFS,
//     marking all 1's as visited, so that we only count an island once.
//
// TC == O(r*c)
// SC == O(r+c) // see alternate solution to reduce this to constant space O(1)
//
// ALTERNATE SOLUTION: In real world, we'd probably want to not change the input matrix.
// This ain't the real world, so we can just mark the matrix elements as 0, to cut down on storage space.
// Give this a try below at some point in the future for some good practice.
func numberOfIslands(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	count := 0 // Count of number of islands

	// deltaRow and deltaCol are for getting next elements
	deltaRow := []int{1, -1, 0, 0}
	deltaCol := []int{0, 0, 1, -1}

	totalRows := len(matrix)
	totalCols := len(matrix[0])

	// visited map tracks which elements we've visited, to avoid an infinite loop
	visited := graph.NewVisited(totalRows, totalCols)

	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalCols; j++ {

			// Already visited this element, so continue
			if visited[i][j] {
				continue
			}

			// Element is a zero, mark it as visited and continue on
			if matrix[i][j] == 0 {
				visited[i][j] = true
				continue
			}

			// Element is a one, enqueue first element in a new island
			q := queue.New(10, state{
				row: i,
				col: j,
			})

			count++ // New island, increment count

			// Found a 1. Run BFS and mark all 1's as visited.
			// Loop over every 1 in the island, marking them as visited
			for !q.Empty() {
				curr := q.DeqHead()
				visited[curr.row][curr.col] = true

				for k := 0; k < len(deltaRow); k++ {
					// TODO: Don't need separate vars for next row and col, can use a new state
					nextRow := curr.row + deltaRow[k]
					nextCol := curr.col + deltaCol[k]

					// TODO: If's are ugly, can be cleaned up
					if nextRow < 0 || nextRow >= totalRows {
						// If next row is out of bounds, continue
						continue
					} else if nextCol < 0 || nextCol >= totalCols {
						// If next column is out of bounds, continue
						continue
					} else if visited[nextRow][nextCol] {
						// If we've already visited this element, continue
						continue
					} else if matrix[nextRow][nextCol] == 0 {
						// If the next element is a zero, mark it visited and continue
						visited[nextRow][nextCol] = true
						continue
					}

					// Next element is a 1 in an island, enqueue it and check next
					q.EnqTail(state{
						row: nextRow,
						col: nextCol,
					})
				}
			}
		}
	}

	return count
}
