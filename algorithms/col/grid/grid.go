package grid

import (
	"fmt"
	"strings"
)

// Cell represents a cell in a grid.
type Cell[T any] struct {
	Row int // The row index of the cell.
	Col int // The column index of the cell.
	Val T   // The value stored in the cell.
}

// Grid represents a grid structure with a specified number of rows and columns.
// Jagged grids are not supported; the number of rows must divide the length of the slice evenly.
type Grid[T any] struct {
	rows  int
	cols  int
	cells []Cell[T]
}

// New creates a new Grid with the specified number of rows and columns.
// The grid is initialized with the zero value of T.
func New[T any](rows int, cols int) *Grid[T] {
	g := &Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: make([]Cell[T], rows*cols),
	}

	for i := 0; i < len(g.cells); i++ {
		g.cells[i].Row = i / rows
		g.cells[i].Col = i % rows
	}

	return g
}

// Initialize sets all the cells in the grid to the specified value.
func (g *Grid[T]) Initialize(rows int, cols int, init T) *Grid[T] {
	for i := 0; i < len(g.cells); i++ {
		g.cells[i].Val = init
	}

	return g
}

// From1D creates a new Grid from a one-dimensional slice.
// Jagged grids are not supported; the number of rows must divide the length of the slice evenly.
func From1D[T any](rows int, x []T) *Grid[T] {
	if len(x)%rows > 0 {
		panic("the rows must be of equal length; len(x)%rows must equal 0")
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

	return &Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

// From2D creates a new Grid from a two-dimensional slice.
// Jagged grids are not supported; the number of rows must divide the length of the slice evenly.
func From2D[T any](x [][]T) *Grid[T] {
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

	return &g
}

func (g *Grid[T]) Vals() []T {
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

func (g *Grid[T]) Get(r, c int) T {
	return g.GetCell(r, c).Val
}

func (g *Grid[T]) Set(r int, c int, x T) *Grid[T] {
	if !g.validRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.validCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	g.cells[g.idx(r, c)].Val = x

	return g
}

func (g *Grid[T]) Row(r int) []T {
	return ctov(g.RowCells(r))
}

func (g *Grid[T]) Col(c int) []T {
	return ctov(g.ColCells(c))
}

func (g *Grid[T]) Neighbors(r int, c int) []T {
	return ctov(g.NeighborCells(r, c))
}

func (g *Grid[T]) LastRow() []T {
	return g.Row(g.rows - 1)
}

func (g *Grid[T]) LastCol() []T {
	return g.Col(g.cols - 1)
}

func (g *Grid[T]) Last() T {
	return g.cells[len(g.cells)-1].Val
}

func (g *Grid[T]) Cells() []Cell[T] {
	return g.cells
}

func (g *Grid[T]) GetCell(r, c int) Cell[T] {
	if !g.validRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.validCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	return g.cells[g.idx(r, c)]
}

func (g *Grid[T]) SetCell(c Cell[T]) *Grid[T] {
	if !g.validRow(c.Row) {
		panic(fmt.Sprintf("row %d is out of range", c.Row))
	} else if !g.validCol(c.Col) {
		panic(fmt.Sprintf("column %d is out of range", c.Col))
	}

	g.cells[g.idx(c.Row, c.Col)] = c

	return g
}

func (g *Grid[T]) RowCells(r int) []Cell[T] {
	if !g.validRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	}

	start, end := g.rowIdx(r)
	return g.cells[start:end]
}

func (g *Grid[T]) ColCells(c int) []Cell[T] {
	if !g.validCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}

	cells := make([]Cell[T], 0, g.cols)

	for i := 0; i < g.rows; i++ {
		cells = append(cells, g.GetCell(i, c))
	}

	return cells
}

func (g *Grid[T]) NeighborCells(r int, c int) []Cell[T] {
	if !g.validRow(r) {
		panic(fmt.Sprintf("row %d is out of range", r))
	} else if !g.validCol(c) {
		panic(fmt.Sprintf("column %d is out of range", c))
	}
	neighbors := make([]Cell[T], 0)

	if r-1 >= 0 {
		neighbors = append(neighbors, g.GetCell(r-1, c))
	}

	if c+1 < g.cols {
		neighbors = append(neighbors, g.GetCell(r, c+1))
	}

	if r+1 < g.rows {
		neighbors = append(neighbors, g.GetCell(r+1, c))
	}

	if c-1 >= 0 {
		neighbors = append(neighbors, g.GetCell(r, c-1))
	}

	return neighbors
}

func (g *Grid[T]) Sub(r1, c1, r2, c2 int) *Grid[T] {
	if !g.validRow(r1) {
		panic(fmt.Sprintf("r1 %d is out of range", r1))
	} else if !g.validRow(r2) {
		panic(fmt.Sprintf("r2 %d is out of range", r2))
	} else if !g.validCol(c1) {
		panic(fmt.Sprintf("c1 %d is out of range", c1))
	} else if !g.validCol(c2) {
		panic(fmt.Sprintf("c2 %d is out of range", c2))
	} else if r1 >= r2 {
		panic("r2 must be greater than r1")
	} else if c1 >= c2 {
		panic("c2 must be greater than c1")
	}

	rows := r2 + 1 - r1
	cols := c2 + 1 - c1

	cells := make([]Cell[T], 0, rows*cols)

	for i := r1; i < r2+1; i++ {
		cells = append(cells, g.RowCells(i)[c1:c2+1]...)
	}

	return &Grid[T]{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func (g *Grid[T]) String() string {
	var sb strings.Builder

	for r := 0; r < g.rows; r++ {
		sb.WriteString(fmt.Sprint(g.Row(r), "\n"))
	}

	return sb.String()
}

func (g *Grid[T]) idx(r int, c int) int {
	return r*g.cols + c
}

func (g *Grid[T]) rowIdx(r int) (int, int) {
	return r * g.cols, r*(g.cols) + g.cols
}

func (g *Grid[T]) validRow(r int) bool {
	return r >= 0 && r < g.rows
}

func (g *Grid[T]) validCol(c int) bool {
	return c >= 0 && c < g.cols
}

func ctov[T any](cells []Cell[T]) []T {
	values := make([]T, 0, len(cells))

	for _, cell := range cells {
		values = append(values, cell.Val)
	}

	return values
}
