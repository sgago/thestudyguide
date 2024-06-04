package knighminimummoves

import (
	"fmt"
	"sgago/thestudyguide/col/vec"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Knight Minimum Moves
https://algo.monster/problems/knight_shortest_path

On an infinitely large chessboard, a knight is located on [0, 0].
A knight in chess can move in eight directions.

Given a destination coordinate [x, y], determine the minimum
number of moves from [0, 0] to [x, y].
*/

func TestKnightMinimumMoves(t *testing.T) {
	actual := knightMinimumMoves(vec.New2D(1, 2))

	fmt.Println(actual)
	assert.Equal(t, 1, actual)
}

func knightMinimumMoves(start vec.V2D[int]) (moves int) {
	if start.Eq(vec.Origin2[int]()) {
		return moves
	}

	visited := make(map[vec.V2D[int]]bool, 10)

	newMoves := getMoves(start, visited)

	fmt.Println(newMoves)

	return moves
}

func getMoves(curr vec.V2D[int], visited map[vec.V2D[int]]bool) (moves []vec.V2D[int]) {
	moves = make([]vec.V2D[int], 0, 8)

	rows := []int{1, 1, 2, 2, -1, -1, -2, -2}
	cols := []int{2, -2, 1, -1, 2, -2, 1, -1}

	for i := 0; i < len(rows); i++ {
		next := vec.New2D(rows[i], cols[i]).Add(&curr)

		if _, ok := visited[next]; !ok {
			moves = append(moves, next)
		}
	}

	return moves
}
