package dsu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	dsu := New[int](4)

	dsu.Add(0)
	dsu.Add(1)
	dsu.Add(2)
	dsu.Add(3)

	fmt.Println(dsu)

	dsu.Union(0, 1)
	dsu.Union(1, 2)
	dsu.Union(2, 3)

	dsu.UnionAdd(5, 1)
	dsu.UnionAdd(6, 2)

	fmt.Println(dsu)

	actual, err := dsu.Find(3)

	fmt.Println(dsu)

	assert.NoError(t, err)
	assert.Equal(t, 0, actual)
}

func TestSame(t *testing.T) {
	dsu := New[int](4)

	dsu.Add(1)
	dsu.UnionAdd(1, 2, 3)
	dsu.Add(4)

	same, _ := dsu.Same(1, 3)
	fmt.Printf("same(%v, %v): %v\n", 1, 3, same)
	assert.True(t, same)

	same, _ = dsu.Same(2, 4)
	fmt.Printf("same(%v, %v): %v\n", 2, 3, same)
	assert.False(t, same)
}

func TestCount(t *testing.T) {
	dsu := New[int](4)

	dsu.Add(1)
	dsu.UnionAdd(1, 2, 3)
	dsu.Add(4)

	fmt.Println(dsu.Len(3))
	count, _ := dsu.Len(3)
	assert.Equal(t, 3, count)

	fmt.Println(dsu.Len(4))
	count, _ = dsu.Len(4)
	assert.Equal(t, 1, count)
}
