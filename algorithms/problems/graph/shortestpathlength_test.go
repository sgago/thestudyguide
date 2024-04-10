package graph

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"
)

type state struct {
	vertex int
	lvl    int
}

func TestShortestPathLength(t *testing.T) {
	from, to := 0, 3

	graph := map[int][]int{
		0: {1, 2},
		1: {0, 2, 3},
		2: {0, 1},
		3: {1},
	}

	visited := map[int]bool{}

	q := queue.New[state](len(graph), state{
		vertex: from,
		lvl:    0,
	})

	for !q.Empty() {
		curr := q.DeqHead()
		visited[curr.vertex] = true

		neighbors := graph[curr.vertex]

		for _, neighbor := range neighbors {

			if neighbor == to {
				fmt.Println(curr.lvl + 1)
				return
			}

			if !visited[neighbor] {
				q.EnqTail(state{
					vertex: neighbor,
					lvl:    curr.lvl + 1,
				})
			}
		}
	}
}
