// Package lru implements a simple LRU (Least Recently Used) cache.
package lru

import (
	"sgago/thestudyguide/col/kvp"
	"sgago/thestudyguide/utils/slices"
)

// Lru represents a simple LRU cache.
type Lru[K comparable, T any] struct {
	s   []kvp.Kvp[K, T] // s is the slice of key-value pairs.
	m   map[K]int       // m is a map holding the index of keys in the slice.
	cap int             // cap is the capacity of the LRU cache.
}

// New creates a new LRU cache with the given capacity.
func New[K comparable, T any](cap int) Lru[K, T] {
	return Lru[K, T]{
		s:   make([]kvp.Kvp[K, T], 0, cap),
		m:   make(map[K]int, cap),
		cap: cap,
	}
}

// Get retrieves the value associated with the given key from the cache.
// It returns the value and a boolean indicating whether the key exists in the cache.
func (lru *Lru[K, T]) Get(key K) (T, bool) {
	idx, ok := lru.m[key]

	if !ok {
		return *new(T), false
	}

	if idx > 0 {
		otherKey := lru.s[idx-1].Key
		otherIdx := lru.m[otherKey]

		lru.m[key], lru.m[otherKey] = lru.m[otherKey], lru.m[key]

		slices.Swap(lru.s, idx, otherIdx)
	}

	return lru.s[idx].Val, true
}

// Put inserts the given key-value pair into the cache.
// If the cache is full, it removes the least recently used item before inserting the new one.
func (lru *Lru[K, T]) Put(key K, val T) {
	if len(lru.s) < lru.cap {
		lru.s = append(lru.s, kvp.New(key, val))
	} else {
		delete(lru.m, lru.s[len(lru.s)-1].Key)
		lru.s[len(lru.s)-1] = kvp.New(key, val)
	}

	lru.m[key] = len(lru.s) - 1
}
