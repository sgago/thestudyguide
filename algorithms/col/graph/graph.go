package graph

func NewVisited(rows, cols int) map[int]map[int]bool {
	visited := make(map[int]map[int]bool, rows)

	for i := 0; i < cols; i++ {
		visited[i] = make(map[int]bool, cols)
	}

	return visited
}
