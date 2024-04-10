// Package kvp provides a simple key-value pair implementation.
package kvp

// Kvp represents a simple key-value pair.
type Kvp[K any, V any] struct {
	Key K // Key is the key of the key-value pair.
	Val V // Val is the value associated with the key.
}

// New creates a new Kvp with the given key and value.
func New[K any, V any](k K, v V) Kvp[K, V] {
	return Kvp[K, V]{
		Key: k,
		Val: v,
	}
}
