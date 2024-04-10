package matrix

import "sgago/thestudyguide/col/dir"

type Matrix[T any] struct {
	elems []T
	rows  int
	cols  int
}

func New[T any](rows, cols int) Matrix[T] {
	m := Matrix[T]{
		elems: make([]T, rows*cols),
		rows:  rows,
		cols:  cols,
	}

	return m
}

func NewInitialize[T any](rows, cols int, init T) Matrix[T] {
	m := Matrix[T]{
		elems: make([]T, rows*cols),
		rows:  rows,
		cols:  cols,
	}

	for i := 0; i < m.Len(); i++ {
		m.elems[i] = init
	}

	return m
}

func (m *Matrix[T]) rowIdx(row int) int {
	return row * m.cols
}

func (m *Matrix[T]) Row(x int) []T {
	start := m.rowIdx(x)
	end := start + m.cols

	return (*m).elems[start:end]
}

func (m *Matrix[T]) Rows() int {
	return m.rows
}

func (m *Matrix[T]) Cols() int {
	return m.cols
}

func (m *Matrix[T]) Len() int {
	return m.rows * m.cols
}

func (m *Matrix[T]) Get(row, col int) T {
	return m.elems[m.rowIdx(row)+col]
}

func (m *Matrix[T]) Set(row, col int, val T) {
	m.elems[m.rowIdx(row)+col] = val
}

type Neighbor[T any] struct {
	row   int
	col   int
	value T
}

func (m *Matrix[T]) Neighbors(row, col int) map[dir.Dir]Neighbor[T] {
	neighbors := make(map[dir.Dir]Neighbor[T])

	// Up
	if row-1 >= 0 {
		neighbors[dir.U] = Neighbor[T]{
			row:   row - 1,
			col:   col,
			value: m.Get(row-1, col),
		}
	}

	// Down
	if row+1 < m.rows {
		neighbors[dir.D] = Neighbor[T]{
			row:   row + 1,
			col:   col,
			value: m.Get(row+1, col),
		}
	}

	// Left
	if col-1 >= 0 {
		neighbors[dir.L] = Neighbor[T]{
			row:   row,
			col:   col - 1,
			value: m.Get(row, col-1),
		}
	}

	// Right
	if col+1 < m.cols {
		neighbors[dir.R] = Neighbor[T]{
			row:   row,
			col:   col + 1,
			value: m.Get(row, col+1),
		}
	}

	return neighbors
}
