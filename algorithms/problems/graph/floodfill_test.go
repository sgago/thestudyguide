package graph

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"
)

/*
	Flood Fill
	https://algo.monster/problems/flood_fill
*/

func TestFloodFill(t *testing.T) {
	row := 2
	col := 2
	replacement := 9

	img := [][]int{
		{0, 1, 3, 4, 1},
		{3, 8, 8, 3, 3},
		{6, 7, 8, 8, 3},
		{2, 2, 8, 9, 1},
		{2, 3, 1, 3, 2},
	}

	fmt.Println("Before:")
	for ri, row := range img {
		fmt.Println(ri, row)
	}

	FloodFillBetter(row, col, replacement, &img)

	fmt.Println("After:")
	for ri, row := range img {
		fmt.Println(ri, row)
	}
}

type coord struct {
	row int
	col int
}

func FloodFill(row int, col int, replacement int, img *[][]int) {
	old := (*img)[row][col] // Replace old with replacement (new) color
	visited := map[int]map[int]bool{}

	width := len(*img)
	height := len((*img)[0])

	// Idea 1: We could maybe just flatten the image and use multiplication to jump among rows and add/sub for cols.

	q := queue.New(8, coord{
		row: row,
		col: col,
	})

	for !q.Empty() {
		curr := q.DeqHead()

		// Reflection 2: Watch out for naming things, this got mega confusing with row vs curr.row vs r.
		// I conflated row and curr.row almost instantly.
		r, c := curr.row, curr.col

		if visited[r] == nil {
			// Reflection 1: Watch maps of maps and slices of slices, they aren't instantiated automagically.
			visited[r] = make(map[int]bool)
		}

		// Improvement 2: Do we even need visited? If we're replacing the color anyway?
		// On changing this to false, we can just rip this out lol.
		// Reflection 3: Think if we need this visited structure when doing these problems or can we do it in place.
		visited[r][c] = true
		(*img)[r][c] = replacement

		// Improvement 1: Solutions have delta rows/cols + for loops that reduce duplication.
		// Reimplement with that.
		if r-1 >= 0 && !visited[r-1][c] && (*img)[r-1][c] == old {
			q.EnqTail(coord{
				row: r - 1,
				col: c,
			})
		}

		if r+1 < height && !visited[r+1][c] && (*img)[r+1][c] == old {
			q.EnqTail(coord{
				row: r + 1,
				col: c,
			})
		}

		if c-1 >= 0 && !visited[r][c-1] && (*img)[r][c-1] == old {
			q.EnqTail(coord{
				row: r,
				col: c - 1,
			})
		}

		if c+1 < width && !visited[r][c+1] && (*img)[r][c+1] == old {
			q.EnqTail(coord{
				row: r,
				col: c + 1,
			})
		}
	}
}

func FloodFillBetter(row int, col int, replacement int, img *[][]int) {
	old := (*img)[row][col] // Replace old with replacement (new) color

	if old == replacement {
		return
	}

	width := len(*img)
	height := len((*img)[0])

	// Use "up" as the start of our arrays, in row/col terms
	// Up = 1,0
	// Down = -1, 0
	// Right = 0, 1
	// Left = 0, -1
	deltaR := []int{1, -1, 0, 0}
	deltaC := []int{0, 0, 1, -1}

	q := queue.New(8, coord{
		row: row,
		col: col,
	})

	for !q.Empty() {
		curr := q.DeqHead()
		r, c := curr.row, curr.col

		(*img)[r][c] = replacement

		for i := 0; i < len(deltaR); i++ {
			nextR := r + deltaR[i]
			nextC := c + deltaC[i]

			if nextR >= 0 &&
				nextR < height &&
				nextC >= 0 &&
				nextC < width &&
				(*img)[nextR][nextC] == old {
				q.EnqTail(coord{
					row: nextR,
					col: nextC,
				})
			}
		}
	}
}
