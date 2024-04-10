package grid

import "fmt"

type Cell[T any] struct {
	Row int
	Col int
	Val T
}

type Grid[T any] struct {
	rows  int
	cols  int
	cells []Cell[T]
}

func New[T any](rows int, cols int) Grid[T] {
	return Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: make([]Cell[T], rows*cols),
	}
}

func NewInit[T any](rows int, cols int, init T) Grid[T] {
	g := New[T](rows, cols)

	for i := 0; i < len(g.cells); i++ {
		g.cells[i].Val = init
	}

	return g
}

func From1D[T any](rows int, x []T) Grid[T] {
	if len(x)%rows > 0 {
		panic("the slices must be of equal length; len(x)%3 must equal 0")
	}

	cols := len(x) / rows

	cells := make([]Cell[T], 0, len(x))

	for i := 0; i < len(x); i++ {
		cells = append(cells, Cell[T]{
			Row: i / rows,
			Col: i % rows,
			Val: x[i],
		})
	}

	return Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func From2D[T any](x [][]T) Grid[T] {
	g := Grid[T]{
		rows: len(x),
		cols: len(x[0]),
	}

	for r, row := range x {
		if len(row) != g.cols {
			panic("each slice must be of equal length")
		}

		for c, value := range row {
			g.cells = append(g.cells, Cell[T]{
				Row: r,
				Col: c,
				Val: value,
			})
		}
	}

	return g
}

func (g *Grid[T]) ValidRow(r int) bool {
	if r < 0 || r >= g.rows {
		return false
	}

	return true
}

func (g *Grid[T]) ValidCol(c int) bool {
	if c < 0 || c >= g.cols {
		return false
	}

	return true
}

func (g *Grid[T]) Cells() []Cell[T] {
	return g.cells
}

func (g *Grid[T]) Values() []T {
	return ctov(g.cells)
}

func (g *Grid[T]) Len() int {
	return len(g.cells)
}

func (g *Grid[T]) Rows() int {
	return g.rows
}

func (g *Grid[T]) Cols() int {
	return g.cols
}

func (g *Grid[T]) Index(r int, c int) int {
	return r*g.cols + c
}

func (g *Grid[T]) RowIndexes(r int) (int, int) {
	return r * g.cols, r*(g.cols) + g.cols
}

func (g *Grid[T]) Get(r, c int) Cell[T] {
	if !g.ValidRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.ValidCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	return g.cells[g.Index(r, c)]
}

func (g *Grid[T]) GetValue(r, c int) T {
	if !g.ValidRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.ValidCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	return g.cells[g.Index(r, c)].Val
}

func (g *Grid[T]) SetValue(r int, c int, x T) {
	if !g.ValidRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.ValidCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	g.cells[g.Index(r, c)].Val = x
}

func (g *Grid[T]) Row(r int) []Cell[T] {
	if !g.ValidRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	}

	start, end := g.RowIndexes(r)
	return g.cells[start:end]
}

func (g *Grid[T]) RowValues(r int) []T {
	return ctov(g.Row(r))
}

func (g *Grid[T]) Col(c int) []Cell[T] {
	if !g.ValidCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	cells := make([]Cell[T], 0, g.cols)

	for i := 0; i < g.rows; i++ {
		cells = append(cells, g.Get(i, c))
	}

	return cells
}

func (g *Grid[T]) ColValues(c int) []T {
	return ctov(g.Col(c))
}

func (g *Grid[T]) Neighbors(r int, c int) []Cell[T] {
	if !g.ValidRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.ValidCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}
	neighbors := make([]Cell[T], 0)

	if r-1 >= 0 {
		neighbors = append(neighbors, g.Get(r-1, c))
	}

	if c+1 < g.cols {
		neighbors = append(neighbors, g.Get(r, c+1))
	}

	if r+1 < g.rows {
		neighbors = append(neighbors, g.Get(r+1, c))
	}

	if c-1 >= 0 {
		neighbors = append(neighbors, g.Get(r, c-1))
	}

	return neighbors
}

func (g *Grid[T]) NeighborValues(r int, c int) []T {
	return ctov(g.Neighbors(r, c))
}

func (g *Grid[T]) Sub(r1, c1, r2, c2 int) Grid[T] {
	if !g.ValidRow(r1) {
		panic(fmt.Sprintf("r1 %d is out of range", r1))
	} else if !g.ValidRow(r2) {
		panic(fmt.Sprintf("r2 %d is out of range", r2))
	} else if !g.ValidCol(c1) {
		panic(fmt.Sprintf("c1 %d is out of range", c1))
	} else if !g.ValidCol(c2) {
		panic(fmt.Sprintf("c2 %d is out of range", c2))
	}

	if r1 >= r2 {
		panic("r2 must be greater than r1")
	} else if c1 >= c2 {
		panic("c2 must be greater than c1")
	}

	rows := r2 + 1 - r1
	cols := c2 + 1 - c1

	cells := make([]Cell[T], 0, rows*cols)

	for i := r1; i < r2+1; i++ {
		cells = append(cells, g.Row(i)[c1:c2+1]...)
	}

	return Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func ctov[T any](cells []Cell[T]) []T {
	values := make([]T, 0, len(cells))

	for _, cell := range cells {
		values = append(values, cell.Val)
	}

	return values
}
