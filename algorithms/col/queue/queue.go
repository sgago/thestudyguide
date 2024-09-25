// Package queue implements a first-in first-out (FIFO) collection.
package queue

import (
	"fmt"

	"sgago/thestudyguide/col/lock"
	"sgago/thestudyguide/errs"
)

// Queue is a a first-in first-out (FIFO) collection.
type Queue[T any] struct {
	lock.RW

	// The elements in this collection.
	v []T
}

// New creates a new Queue and initializes it with the given values.
func New[T any](cap int, vals ...T) *Queue[T] {
	cap = max(cap, len(vals))

	return &Queue[T]{
		v: append(make([]T, 0, cap), vals...),
	}
}

// Slice returns the internal slice backing this collection.
func (q *Queue[T]) Slice() []T {
	defer q.RW.Read()()

	return q.v
}

// String returns this collection represented as a string.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", q.Slice())
}

// Len returns the number of elements in this Queue.
func (q *Queue[T]) Len() int {
	defer q.RW.Read()()

	return len(q.v)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

// Cap returns the maximum number of elements this Queue can have before being re-sized.
func (q *Queue[T]) Cap() int {
	defer q.RW.Read()()

	return cap(q.v)
}

// TryDeq attempts to pop a value from the start of the Queue and returns an error if the Queue is empty.
func (q *Queue[T]) TryDeq() (T, error) {
	defer q.RW.Write()()

	if len(q.v) > 0 {
		result := q.v[0]
		q.v = q.v[:len(q.v)-1]
		return result, nil
	}

	return *new(T), errs.Empty
}

// TryPeek attempts to peek at the start of the Queue and returns an error if the Queue is empty.
func (q *Queue[T]) TryPeek() (T, error) {
	defer q.RW.Read()()

	if len(q.v) > 0 {
		return q.v[0], nil
	}

	return *new(T), errs.Empty
}

func (q *Queue[T]) EnqHead(vals ...T) {
	defer q.RW.Write()()

	q.v = append(vals, q.v...)
}

func (q *Queue[T]) EnqTail(vals ...T) {
	defer q.RW.Write()()

	q.v = append(q.v, vals...)
}

func (q *Queue[T]) DeqHead() T {
	defer q.RW.Write()()

	result := q.v[0]
	q.v = q.v[1:q.Len()]
	return result
}

func (q *Queue[T]) TryDeqHead() (T, error) {
	defer q.RW.Write()()

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[0]
	q.v = q.v[1:q.Len()]
	return result, nil
}

func (q *Queue[T]) DeqTail() T {
	defer q.RW.Write()()

	result := q.v[len(q.v)-1]
	q.v = q.v[0 : q.Len()-1]
	return result
}

func (q *Queue[T]) TryDeqTail() (T, error) {
	defer q.RW.Write()()

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[len(q.v)-1]
	q.v = q.v[0 : q.Len()-1]
	return result, nil
}

func (q *Queue[T]) PeekHead() T {
	defer q.RW.Read()()

	result := q.v[0]
	return result
}

func (q *Queue[T]) TryPeekHead() (T, error) {
	defer q.RW.Read()()

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[0]
	return result, nil
}

func (q *Queue[T]) PeekTail() T {
	defer q.RW.Read()()

	result := q.v[len(q.v)-1]
	return result
}

func (q *Queue[T]) TryPeekTail() (T, error) {
	defer q.RW.Read()()

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[len(q.v)-1]
	return result, nil
}
