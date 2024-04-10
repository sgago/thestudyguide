// Package dsu implements a ranked disjoint set union (DSU).
package dsu

import (
	"sgago/thestudyguide/errs"
)

// Dsu represents a Disjoint Set Union data structure.
//
//   - Disjoint means that sets will not have shared elements.
//
//   - Set is a collection with all unique or distinct elements.
//
//   - Union means the operation of merging or combining the disjointed sets.
type Dsu[T comparable] struct {
	// A map of each element to their parent element
	// a.k.a. to their root element or set ID.
	id map[T]T

	// A map with the rank of each element for efficient merging.
	rank map[T]int

	// A map with the count of elements in the set.
	count map[T]int
}

func New[T comparable](cap int) *Dsu[T] {
	return &Dsu[T]{
		id:    make(map[T]T, cap),
		rank:  make(map[T]int, cap),
		count: make(map[T]int, cap),
	}
}

// Contains returns true if this Dsu contains x;
// otherwise, false.
func (dsu *Dsu[T]) Contains(x T) bool {
	_, ok := dsu.id[x]
	return ok
}

// Same returns true if elements x and y are in the same
// set; otherwise, false.
func (dsu *Dsu[T]) Same(x, y T) (bool, error) {
	rootX, err := dsu.Root(x)

	if err != nil {
		return false, err
	}

	rootY, err := dsu.Root(y)

	if err != nil {
		return false, err
	}

	return rootX == rootY, nil
}

// Adds elements to the DSU. Each element will be in its own
// disjoint set.
func (dsu *Dsu[T]) Add(elems ...T) {
	for _, elem := range elems {
		dsu.id[elem] = elem
		dsu.rank[elem] = 0
		dsu.count[elem] = 1
	}
}

// Root finds the parent ID, also know as the set ID, of the target.
func (dsu *Dsu[T]) Root(target T) (T, error) {
	parent, ok := dsu.id[target]

	if !ok {
		bad := new(T)
		return *bad, errs.NotFound
	}

	if parent != target {
		// If x is not it's own parent,
		// make a recursive call to find x's parent.
		newParent, err := dsu.Root(dsu.id[target])

		if err != nil {
			dsu.id[target] = newParent
		}
	}

	return dsu.id[target], nil
}

// Union merges elems into the set that target is in.
func (dsu *Dsu[T]) Union(target T, elems ...T) error {
	rootX, err := dsu.Root(target)

	if err != nil {
		return err
	}

	rankX := dsu.rank[rootX]

	for _, elem := range elems {
		rootY, err := dsu.Root(elem)

		if err != nil {
			return err
		}

		rankY := dsu.rank[rootY]

		if rankX < rankY {
			dsu.id[rootX] = rootY
			dsu.count[rootY] += dsu.count[rootX]
		} else {
			dsu.id[rootY] = rootX
			dsu.count[rootX] += dsu.count[rootY]

			if rankX == rankY {
				dsu.rank[rootX]++
			}
		}
	}

	return nil
}

// Returns the length of the set containing x.
func (dsu *Dsu[T]) Len(x T) (int, error) {
	rootX, err := dsu.Root(x)

	if err != nil {
		return -1, err
	}

	return dsu.count[rootX], nil
}

// UnionAdd adds elements and unions them all to target.
func (dsu *Dsu[T]) UnionAdd(existing T, new ...T) error {
	dsu.Add(new...)
	return dsu.Union(existing, new...)
}
