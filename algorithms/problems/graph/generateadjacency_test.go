package graph

import (
	"fmt"
	"testing"
)

func TestGenerateAdjacency(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	adj := map[int][]int{}

	for rowIdx, row := range matrix {
		for colIdx, val := range row {

			adj[val] = make([]int, 0)

			if colIdx-1 >= 0 {
				adj[val] = append(adj[val], matrix[rowIdx][colIdx-1])
			}

			if colIdx+1 < len(row) {
				adj[val] = append(adj[val], matrix[rowIdx][colIdx+1])
			}

			if rowIdx-1 >= 0 {
				adj[val] = append(adj[val], matrix[rowIdx-1][colIdx])
			}

			if rowIdx+1 < len(matrix) {
				adj[val] = append(adj[val], matrix[rowIdx+1][colIdx])
			}
		}
	}

	fmt.Println(adj)
}
