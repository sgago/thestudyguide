package matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	m := New[int](3, 3)

	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			m.Set(r, c, r*c)
		}
	}

	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			val := m.Get(r, c)
			assert.Equal(t, r*c, val)
		}
	}
}

func TestNeighbors(t *testing.T) {
	m := New[int](3, 3)

	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			m.Set(r, c, r*c)
		}
	}

	n := m.Neighbors(1, 1)

	fmt.Println(n)
}
