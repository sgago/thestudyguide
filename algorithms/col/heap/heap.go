package heap

import (
	"cmp"
	"fmt"

	sliceutil "sgago/thestudyguide/utils/slices"
)

type Heap[T any] struct {
	v    []*T
	less func(i, j T) bool
}

func New[T cmp.Ordered](cap int, vals ...T) *Heap[T] {
	return NewFunc[T](cap, func(i, j T) bool { return i < j }, vals...)
}

// New creates a new min heap using the less function.
func NewFunc[T any](cap int, less func(i, j T) bool, vals ...T) *Heap[T] {
	cap = max(cap, len(vals))

	h := &Heap[T]{
		v:    make([]*T, 0, cap),
		less: less,
	}

	h.Push(vals...)

	return h
}

// Slice returns the underlying slice backing this collection.
func (h *Heap[T]) Slice() []*T {
	return h.v
}

// String returns the underlying slice backing this collection represented as a string.
func (h *Heap[T]) String() string {
	return fmt.Sprintf("%v", h.v)
}

// Len returns the number of elements in this collection.
func (h *Heap[T]) Len() int {
	return len(h.v)
}

// Empty returns true if this collection is empty; otherwise, false.
func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

// Cap returns the capacity of underlying slice backing this collection.
func (h *Heap[T]) Cap() int {
	return cap(h.v)
}

func (h *Heap[T]) Clear() {
	clear(h.v)
}

func (h *Heap[T]) Peek() T {
	return *h.v[0]
}

func (h *Heap[T]) leftIdx(idx int) int {
	return 2*idx + 1
}

func (h *Heap[T]) rightIdx(idx int) int {
	return 2*idx + 2
}

func (h *Heap[T]) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap[T]) safeIdx(idx int) bool {
	return h.Len() > idx
}

func (h *Heap[T]) Push(vals ...T) {
	for i := 0; i < len(vals); i++ {
		h.v = append(h.v, &vals[i])
		h.bubbleUp(h.Len() - 1)
	}
}

func (h *Heap[T]) bubbleUp(idx int) {
	if idx <= 0 {
		return
	}

	parent := h.parentIdx(idx)

	if !h.less(*h.v[parent], *h.v[idx]) {
		sliceutil.Swap(h.v, parent, idx)
	}

	if parent > 0 {
		h.bubbleUp(parent)
	}
}

func (h *Heap[T]) Pop() T {
	val := h.v[0]

	h.v[0] = h.v[len(h.v)-1] // Copy last element to top
	h.v = h.v[:len(h.v)-1]   // Drop the last element

	h.bubbleDown(0) // Bubble top down to bottom

	return *val
}

func (h *Heap[T]) bubbleDown(idx int) {
	left, right := h.leftIdx(idx), h.rightIdx(idx)
	smallest := left

	if !h.safeIdx(left) {
		return
	}

	if h.safeIdx(right) && h.v[right] != nil && !h.less(*h.v[left], *h.v[right]) {
		smallest = right
	}

	if h.v[smallest] != nil && h.less(*h.v[smallest], *h.v[idx]) {
		sliceutil.Swap(h.v, smallest, idx)
		h.bubbleDown(smallest)
	}
}
