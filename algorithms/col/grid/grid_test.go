package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	g := New[int](3, 3)

	for i := 0; i < 3*3; i++ {
		r, c := i/3, i%3

		g.SetValue(r, c, i)
		val := g.GetValue(r, c)

		assert.Equal(t, i, val)
	}
}

func TestFrom2D(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)

	for i := 0; i < 3*3; i++ {
		r, c := i/3, i%3
		val := g.GetValue(r, c)
		assert.Equal(t, i, val)
	}
}

func TestFrom1D(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	g := From1D(3, s)

	for i := 0; i < 3*3; i++ {
		r, c := i/3, i%3
		val := g.GetValue(r, c)
		assert.Equal(t, i, val)
	}
}

func TestSub(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	sub := g.Sub(1, 1, 2, 2)

	assert.Equal(t, []int{4, 5, 7, 8}, sub.Values())
}

func TestRow(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	col := g.RowValues(1)

	assert.Equal(t, []int{3, 4, 5}, col)
}

func TestCol(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	col := g.ColValues(1)

	assert.Equal(t, []int{1, 4, 7}, col)
}

func TestMiddleNeighbors(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	neighbors := g.NeighborValues(1, 1)

	assert.Equal(t, []int{1, 5, 7, 3}, neighbors)
}

func TestUpperLeftCornerNeighbors(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	neighbors := g.NeighborValues(0, 0)

	assert.Equal(t, []int{1, 3}, neighbors)
}

func TestBottomRightCornerNeighbors(t *testing.T) {
	s := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	g := From2D(s)
	neighbors := g.NeighborValues(2, 2)

	assert.Equal(t, []int{5, 7}, neighbors)
}
