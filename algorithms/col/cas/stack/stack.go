// Package cas/stack implements a compare-and-swap (CAS) first-in
// last-out (FILO) collection that supports concurrent operations.
// Reference: https://www.geeksforgeeks.org/lock-free-stack-using-java/
package stack

import (
	"fmt"
	"sync/atomic"

	"sgago/thestudyguide/errs"
)

// node represents a node in the CAS stack, containing a value and a pointer to the next node.
type node[T any] struct {
	// The node's value.
	val T

	// A pointer to the next node.
	next atomic.Pointer[node[T]]
}

// Stack is a compare-and-swap (CAS) stack that supports concurrent operations.
type Stack[T any] struct {
	// A pointer to the top (head) node of the stack.
	head atomic.Pointer[node[T]]

	// The number of elements in the stack.
	len atomic.Int32
}

// New creates a new CasStack and initializes it with the given values.
func New[T any](vals ...T) *Stack[T] {
	s := &Stack[T]{}
	s.Push(vals...)
	return s
}

// Len returns the current length of the stack.
func (s *Stack[T]) Len() int {
	return int(s.len.Load())
}

func (s *Stack[T]) Empty() bool {
	return s.Len() > 0
}

// Push adds values to the top of the stack using CAS operations for concurrency safety.
func (s *Stack[T]) Push(vals ...T) {
	for _, v := range vals {
		new := &node[T]{
			val:  v,
			next: s.head,
		}

		for !s.head.CompareAndSwap(s.head.Load(), new) {
			new.next = s.head
		}

		s.len.Add(1)
	}
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack[T]) Peek() T {
	return s.head.Load().val
}

// TryPeek attempts to peek at the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPeek() (T, error) {
	h := s.head.Load()

	if h == nil {
		return *new(T), errs.Empty
	}

	s.len.Add(-1)

	return h.val, nil
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack[T]) Pop() T {
	curr := s.head.Load()
	// This is where the magic happens.
	// If s.head is the same as curr, then we swap.
	// If s.head is different because another thread modified it, we spin: reload curr and try it again.
	for !s.head.CompareAndSwap(s.head.Load(), curr.next.Load()) {
		curr = s.head.Load()
	}

	s.len.Add(-1)

	return curr.val
}

// TryPop attempts to pop a value from the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPop() (T, error) {
	curr := s.head.Load()

	for curr != nil && !s.head.CompareAndSwap(curr, curr.next.Load()) {
		curr = s.head.Load()
	}

	if curr == nil {
		return *new(T), errs.Empty
	}

	s.len.Add(-1)

	return curr.val, nil
}

func (s *Stack[T]) Slice() []T {
	results := make([]T, 0, s.Len())

	curr := s.head.Load()
	for curr != nil {
		results = append(results, curr.val)
		curr = curr.next.Load()
	}

	return results
}

func (s *Stack[T]) String() string {
	results := s.Slice()
	return fmt.Sprintf("%v", results)
}
