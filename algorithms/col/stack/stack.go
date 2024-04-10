// Package stack implements a first-in last-out (FILO) collection.
package stack

import (
	"fmt"

	"sgago/thestudyguide/col/conc"
	"sgago/thestudyguide/errs"
	sliceutil "sgago/thestudyguide/utils/slices"
)

// Stack is a a first-in last-out (FILO) collection.
type Stack[T any] struct {
	conc.RWLocker

	elems []T
	head  int
}

// New creates a new Stack and initializes it with the given values.
func New[T any](cap int, vals ...T) *Stack[T] {
	s := &Stack[T]{
		elems: make([]T, cap),
		head:  -1,
	}

	s.Push(vals...)

	return s
}

// Slice returns the internal slice backing this collection.
func (s *Stack[T]) Slice() []T {
	defer s.RWLocker.Write()()

	return s.elems[:s.head+1]
}

// String returns this collection represented as a string.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

// Len returns the number of elements in this Stack.
func (s *Stack[T]) Len() int {
	defer s.RWLocker.Read()()

	return s.head + 1
}

// Cap returns current capacity of the underlying slice backing this stack.
func (s *Stack[T]) Cap() int {
	defer s.RWLocker.Read()()

	return cap(s.elems)
}

// Empty returns true if the stack is empty; otherwise, false.
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Resize changes the length of the underlying slice by best-effort.
// It will not delete stack elements.
func (s *Stack[T]) Resize(cap int) {
	defer s.RWLocker.Write()()

	s.elems = s.elems[:max(len(s.elems), cap)]
}

// Push adds values to the top of the Stack.
func (s *Stack[T]) Push(vals ...T) {
	if len(vals) == 0 {
		return
	}

	defer s.RWLocker.Write()()

	newHead := s.head + len(vals)

	// Does our slice have enough space?
	// If not, double its length to avoid calling append frequently because it's slow.
	if newHead >= len(s.elems) {
		s.elems = append(s.elems, make([]T, len(s.elems)+len(vals))...)
	}

	// Instead of pushing each element in a tight loop, reverse the elements
	// and copy them over to save time.
	sliceutil.Reverse(vals)
	copy(s.elems[s.head+1:], vals)

	s.head = newHead
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack[T]) Pop() T {
	defer s.RWLocker.Write()()

	if s.head == -1 {
		panic("the stack is empty")
	}

	s.head--

	return s.elems[s.head+1]
}

// TryPop attempts to pop a value from the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPop() (T, error) {
	defer s.RWLocker.Write()()

	if s.head >= 0 {
		s.head--
		return s.elems[s.head+1], nil
	}

	return *new(T), errs.Empty
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack[T]) Peek() T {
	defer s.RWLocker.Read()()

	if s.head == -1 {
		panic("the stack is empty")
	}

	return s.elems[s.head]
}

// TryPeek attempts to peek at the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPeek() (T, error) {
	defer s.RWLocker.Read()()

	if s.head >= 0 {
		return s.elems[s.head], nil
	}

	return *new(T), errs.Empty
}
