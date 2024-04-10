// Package cas/queue implements a compare-and-swap (CAS) first-in
// first-out (FIFO) collection that supports concurrent operations.
// Reference: https://www.geeksforgeeks.org/lock-free-stack-using-java/
package queue

import (
	"fmt"
	"sync/atomic"
)

// node represents a node in the CAS queue, containing a value and a pointer to the next node.
type node[T any] struct {
	// The node's value.
	val T

	// A pointer to the next node in the queue.
	next atomic.Pointer[node[T]]

	// A pointer to the previous node in queue.
	prev atomic.Pointer[node[T]]
}

// Queue is a compare-and-swap (CAS) queue that supports concurrent operations.
type Queue[T any] struct {
	// A pointer to the first (head) node of the queue.
	head atomic.Pointer[node[T]]

	// A pointer to the last (tail) node of the queue.
	tail atomic.Pointer[node[T]]

	// The number of elements in the queue.
	len atomic.Int32
}

func New[T any](vals ...T) *Queue[T] {
	q := &Queue[T]{}
	q.EnqTail(vals...)
	return q
}

// Len returns the current length of the stack.
func (q *Queue[T]) Len() int {
	return int(q.len.Load())
}

func (q *Queue[T]) EnqHead(vals ...T) {
	for _, val := range vals {
		new := &node[T]{
			val:  val,
			next: q.head,
		}

		for !q.head.CompareAndSwap(q.head.Load(), new) {
			new.next = q.head
		}

		for q.tail.Load() == nil && !q.tail.CompareAndSwap(q.tail.Load(), q.head.Load()) {
			// CAS failed, try again
		}

		if new.next.Load() != nil {
			oldVal := new.next.Load()
			oldVal.prev = q.head
			for !new.next.CompareAndSwap(new.next.Load(), oldVal) {
				oldVal.prev = q.head
			}
		}

		q.len.Add(1)
	}
}

// Pop removes and returns the value at the top of the stack.
func (q *Queue[T]) DeqHead() T {
	result := q.head.Load()

	// This is where the magic happens.
	// If q.head is the same as curr, then we swap.
	// If q.head is different because another thread modified it, we spin: reload curr and try it again.
	for !q.head.CompareAndSwap(result, result.next.Load()) {
		result = q.head.Load()
	}

	q.len.Add(-1)

	return result.val
}

func (q *Queue[T]) EnqTail(vals ...T) {
	for _, val := range vals {
		new := &node[T]{
			val:  val,
			prev: q.tail,
		}

		for !q.tail.CompareAndSwap(q.tail.Load(), new) {
			new.prev = q.tail
		}

		for q.head.Load() == nil && !q.head.CompareAndSwap(q.head.Load(), q.tail.Load()) {
			// CAS failed, try again
		}

		if new.prev.Load() != nil {
			oldVal := new.prev.Load()
			oldVal.next = q.tail
			for !new.prev.CompareAndSwap(new.prev.Load(), oldVal) {
				oldVal.next = q.tail
			}
		}

		q.len.Add(1)
	}
}

func (q *Queue[T]) DeqTail() T {
	result := q.tail.Load()

	for !q.tail.CompareAndSwap(result, result.prev.Load()) {
		result = q.tail.Load()
	}

	q.len.Add(-1)

	return result.val
}

func (q *Queue[T]) PeekHead() T {
	return q.head.Load().val
}

func (q *Queue[T]) PeekTail() T {
	return q.tail.Load().val
}

func (q *Queue[T]) Slice() []T {
	results := make([]T, 0, q.Len())

	curr := q.head.Load()
	for curr != nil {
		results = append(results, curr.val)
		curr = curr.next.Load()
	}

	return results
}

func (q *Queue[T]) String() string {
	results := q.Slice()
	return fmt.Sprintf("%v", results)
}
