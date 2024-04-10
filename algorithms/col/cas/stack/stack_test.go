package stack

// Reference: https://www.geeksforgeeks.org/lock-free-stack-using-java/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack[int]{}

	s.Push(123)
	result := s.Pop()

	fmt.Printf("%v", result)
}

func TestStackPushConcurrent(t *testing.T) {
	const stackSize = 1000

	s := Stack[int]{}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackSize)
	for i := 0; i < stackSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackSize; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			s.Push(n)
			m[n].Add(1)
		}(i)
	}

	wg.Wait()

	// Verify that all values from [0 to stackSize) were pushed exactly once
	for i := 0; i < stackSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was pushed %d times", i, count)
		}
	}
}

func TestStackPopConcurrent(t *testing.T) {
	const stackSize = 1000

	s := Stack[int]{}

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < stackSize; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackSize)
	for i := 0; i < stackSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()

	// Verify that all values from [0 to stackSize) were popped exactly once
	for i := 0; i < stackSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was popped %d times", i, count)
		}
	}
}
