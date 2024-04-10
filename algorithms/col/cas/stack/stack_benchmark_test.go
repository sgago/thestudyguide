package stack

import (
	"sync"
	"sync/atomic"
	"testing"

	normal "sgago/thestudyguide/col/stack"
)

const stackLen = 10_000

var nonSyncStack *normal.Stack[int]
var syncStack *normal.Stack[int]

var casStack *Stack[int]

func init() {
	nonSyncStack = normal.New[int](stackLen)

	syncStack = normal.New[int](stackLen)
	syncStack.Synchronize(true)

	casStack = New[int]()
}

func BenchmarkNonSyncStackPushPop(b *testing.B) {
	s := nonSyncStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkNonSyncStackStringify(b *testing.B) {
	s := nonSyncStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkSyncStackPushPop(b *testing.B) {
	s := syncStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkSyncStackStringify(b *testing.B) {
	s := syncStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkSyncStackPushConcurrent(b *testing.B) {
	s := syncStack

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackLen)
	for i := 0; i < stackLen; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackLen; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			s.Push(n)
			m[n].Add(1)
		}(i)
	}

	wg.Wait()

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkSyncStackPopConcurrent(b *testing.B) {
	s := syncStack

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackLen)
	for i := 0; i < stackLen; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackLen; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()
}

func BenchmarkCasStackPushPop(b *testing.B) {
	s := casStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkCasStackStringify(b *testing.B) {
	s := casStack

	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkCasStackPopConcurrent(b *testing.B) {
	s := casStack

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < stackLen; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackLen)
	for i := 0; i < stackLen; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackLen; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()
}

func BenchmarkCasStackPushConcurrent(b *testing.B) {
	s := casStack

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackLen)
	for i := 0; i < stackLen; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackLen; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			s.Push(n)
			m[n].Add(1)
		}(i)
	}

	wg.Wait()

	for !s.Empty() {
		s.Pop()
	}
}
