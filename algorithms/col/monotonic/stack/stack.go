package stack

import (
	"cmp"
	"fmt"

	"sgago/thestudyguide/col/lock"
)

type Stack[T any] struct {
	lock.RW

	v    []T
	idx  int
	less func(a, b T) bool
}

func New[T cmp.Ordered](cap int) *Stack[T] {
	s := &Stack[T]{
		v:    make([]T, cap),
		idx:  -1,
		less: func(a, b T) bool { return a < b },
	}

	return s
}

func NewFunc[T any](cap int, less func(a, b T) bool) *Stack[T] {
	s := &Stack[T]{
		v:   make([]T, cap),
		idx: -1,
	}

	return s
}

func (s *Stack[T]) Push(val T) []T {
	defer s.RW.Write()()

	if s.idx < 0 || !s.less(val, s.v[s.idx]) {
		s.idx++

		if s.idx >= len(s.v) {
			s.v = append(s.v, make([]T, len(s.v))...)
		}

		s.v[s.idx] = val

		return make([]T, 0)
	}

	prevIdx := s.idx
	left, right := 0, s.idx

	for left <= right {
		mid := left + (right-left)/2

		if s.less(s.v[mid], val) {
			left = mid + 1
		} else {
			s.idx = mid
			right = mid - 1
		}
	}

	result := make([]T, prevIdx+1-s.idx)
	copy(result, s.v[s.idx:prevIdx+1])

	s.v[s.idx] = val

	return result
}

func (s *Stack[T]) Peek() T {
	defer s.RW.Read()()
	return s.v[s.idx]
}

func (s *Stack[T]) Pop() T {
	defer s.RW.Write()()
	result := s.v[s.idx]
	s.idx--
	return result
}

func (s *Stack[T]) Slice() []T {
	defer s.RW.Write()()
	return s.v[:s.idx+1]
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

func (s *Stack[T]) Len() int {
	defer s.RW.Read()()
	return s.idx + 1
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Resize(cap int) {
	defer s.RW.Write()()
	s.v = s.v[:max(len(s.v), cap)]
}
