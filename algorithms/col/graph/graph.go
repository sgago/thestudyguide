package graph

// NewVisited creates a new visited map with the specified number of rows and columns.
// It is used to keep track of visited nodes in a graph traversal.
func NewVisited(rows, cols int) map[int]map[int]bool {
	visited := make(map[int]map[int]bool, rows)

	for i := 0; i < cols; i++ {
		visited[i] = make(map[int]bool, cols)
	}

	return visited
}
