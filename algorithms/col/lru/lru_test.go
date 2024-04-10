package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {
	lru := New[int, string](3)

	// Add some items to the cache
	lru.Put(1, "one")
	lru.Put(2, "two")
	lru.Put(3, "three")

	// Test Get
	val, found := lru.Get(1)
	assert.True(t, found)
	assert.Equal(t, "one", val)

	// Test Put when the cache is full
	lru.Put(4, "four")
	_, found = lru.Get(3) // 2 should have been removed due to being least recently used
	assert.False(t, found)

	// Test Get for non-existent key
	_, found = lru.Get(5)
	assert.False(t, found)
}
