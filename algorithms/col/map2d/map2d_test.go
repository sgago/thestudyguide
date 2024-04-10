package map2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap2dGet_KeyValueExists(t *testing.T) {
	expected := 123

	m2d := New[int, int](3)
	actual, ok := m2d.Set(1, 1, expected).Get(1, 1)

	assert.True(t, ok)
	assert.Equal(t, expected, actual)
}

func TestMap2dGet_KeyValueDoesNotExists(t *testing.T) {
	expected := 123

	m2d := New[int, int](3)
	_, ok := m2d.Set(1, 1, expected).Get(2, 2)

	assert.False(t, ok)
}
