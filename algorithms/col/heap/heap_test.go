package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := NewFunc[int](3, func(i, j int) bool { return i < j })
	h.Push(3)
	h.Push(2)
	h.Push(1)
	h.Push(6)
	h.Push(5)
	h.Push(4)

	actual := h.Pop()
	assert.Equal(t, 1, actual)

	actual = h.Pop()
	assert.Equal(t, 2, actual)

	actual = h.Pop()
	assert.Equal(t, 3, actual)

	actual = h.Pop()
	assert.Equal(t, 4, actual)

	actual = h.Pop()
	assert.Equal(t, 5, actual)

	actual = h.Pop()
	assert.Equal(t, 6, actual)
}
