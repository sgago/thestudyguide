// Package dsu implements a ranked disjoint set union (DSU).
package dsu

import (
	"sgago/thestudyguide/errs"
)

// Dsu represents a Disjoint Set Union data structure.
//   - Disjoint means that sets will not have shared elements.
//   - Set is a collection with all unique or distinct elements.
//   - Union means the operation of merging or combining the disjointed sets.
type Dsu[T comparable] struct {
	// A map of each element to their parent element
	// a.k.a. to their root element or set ID.
	// The key is the element itself, the value is its parent.
	id map[T]T

	// A map with the rank of each element for efficient merging.
	// This ensures that the depth of the DSU tree is as flat as possible.
	// It is an upper bounds on the height of the tree.
	// This is what gets us an O(alpha(n)) or O(α(n)) time complexity.
	// alpha(n) is very small, like O(4) small which ~= O(1) ~= which is negligible.
	rank map[T]int

	// A map with the count of elements in the set.
	count map[T]int
}

// New instantiates a Disjoint Set Union (DSU) with a given capacity.
func New[T comparable](cap int) *Dsu[T] {
	return &Dsu[T]{
		id:    make(map[T]T, cap),
		rank:  make(map[T]int, cap),
		count: make(map[T]int, cap),
	}
}

// Contains returns true if this Dsu contains x; otherwise, false.
func (dsu *Dsu[T]) Contains(x T) bool {
	_, ok := dsu.id[x]
	return ok
}

// Same returns true if elements x and y are in the same
// set; otherwise, false.
func (dsu *Dsu[T]) Same(x, y T) (bool, error) {
	rootX, err := dsu.Find(x)

	if err != nil {
		return false, err
	}

	rootY, err := dsu.Find(y)

	if err != nil {
		return false, err
	}

	return rootX == rootY, nil
}

// Adds elements to the DSU. Each element will be in its own set.
// This is, it will not be union'ed to anything else.
func (dsu *Dsu[T]) Add(elems ...T) {
	for _, elem := range elems {
		if !dsu.Contains(elem) {
			dsu.id[elem] = elem
			dsu.rank[elem] = 0
			dsu.count[elem] = 1
		}
	}
}

// Find gets the parent ID, also know as the set ID, of the target x.
func (dsu *Dsu[T]) Find(x T) (T, error) {
	parent, ok := dsu.id[x]
	if !ok {
		bad := new(T)
		return *bad, errs.NotFound
	}

	if parent != x {
		// If x is not it's own parent,
		// make a recursive call to find x's parent.
		newParent, err := dsu.Find(dsu.id[x])

		if err != nil {
			dsu.id[x] = newParent
		}

		dsu.id[x] = newParent
		return newParent, nil
	}

	return parent, nil
}

// Union merges elems into the set that target is in.
func (dsu *Dsu[T]) Union(target T, elems ...T) error {
	rootX, err := dsu.Find(target)
	if err != nil {
		return err
	}

	for _, elem := range elems {
		rootY, err := dsu.Find(elem)
		if err != nil {
			return err
		}

		if rootX != rootY {
			if dsu.rank[rootX] > dsu.rank[rootY] {
				// rootY has the smaller tree (rank), so it will
				// be made a subtree of rootX.
				dsu.id[rootY] = rootX
				dsu.count[rootX] += dsu.count[rootY]
			} else if dsu.rank[rootX] < dsu.rank[rootY] {
				// rootX has the smaller tree (rank), so it will
				// be made a subtree of rootY.
				dsu.id[rootX] = rootY
				dsu.count[rootY] += dsu.count[rootX]
			} else {
				// rootX and rootY have same rank.
				// Arbitrarily choose rootY to be a subtree of rootX.
				dsu.id[rootY] = rootX
				dsu.count[rootX] += dsu.count[rootY]
				dsu.rank[rootX]++
			}
		}
	}

	return nil
}

// Len returns the length of the set containing x.
func (dsu *Dsu[T]) Len(x T) (int, error) {
	rootX, err := dsu.Find(x)

	if err != nil {
		return -1, err
	}

	return dsu.count[rootX], nil
}

// UnionAdd adds elements and unions them all to an existing target.
func (dsu *Dsu[T]) UnionAdd(existing T, new ...T) error {
	dsu.Add(new...)
	return dsu.Union(existing, new...)
}

// Set gets all elements in the same set as x.
func (dsu *Dsu[T]) Set(x T) ([]T, error) {
	set := make([]T, 0)

	for child := range dsu.id {
		if ok, err := dsu.Same(x, child); err != nil {
			return nil, err
		} else if ok {
			set = append(set, child)
		}
	}

	return set, nil
}

// Roots gets all the unique root Ids, the parents or set Ids, of the individual sets.
func (dsu *Dsu[T]) Roots() []T {
	parents := make(map[T]T)

	for i := range dsu.id {
		parent, _ := dsu.Find(i)
		parents[parent] = parent
	}

	uniqueParents := make([]T, 0)

	for _, parent := range parents {
		uniqueParents = append(uniqueParents, parent)
	}

	return uniqueParents
}

// Sets get all of the disjointed sets.
func (dsu *Dsu[T]) Sets() [][]T {
	roots := dsu.Roots()

	sets := make([][]T, len(roots))

	for i, root := range roots {
		set, _ := dsu.Set(root)
		sets[i] = set
	}

	return sets
}
