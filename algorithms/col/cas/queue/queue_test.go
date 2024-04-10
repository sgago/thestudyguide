package queue

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueHead1(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)

	result := q.DeqHead()
	assert.Equal(t, 123, result)
}

func TestQueueHead2(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)
	q.EnqHead(456)

	result := q.DeqHead()
	assert.Equal(t, 456, result)

	result = q.DeqHead()
	assert.Equal(t, 123, result)
}

func TestNextPointersFromEnqHead(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)
	q.EnqHead(456)
	q.EnqHead(789)

	ptr := q.head
	actual := ptr.Load()
	assert.Equal(t, 789, actual.val)

	ptr = actual.next
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 456)

	ptr = actual.next
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 123)
}

func TestPrevPointersFromEnqHead(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)
	assert.Equal(t, q.head, q.tail)
	oldHead := q.head
	oldTail := q.tail

	q.EnqHead(456)
	assert.NotEqual(t, oldHead, q.head)
	assert.Equal(t, oldTail, q.tail)
	oldHead = q.head
	oldTail = q.tail

	q.EnqHead(789)
	assert.NotEqual(t, oldHead, q.head)
	assert.Equal(t, oldTail, q.tail)

	ptr := q.tail
	actual := ptr.Load()
	assert.Equal(t, 123, actual.val)

	ptr = actual.prev
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 456)

	ptr = actual.prev
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 789)
}

func TestQueueTail1(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)

	result := q.DeqTail()
	assert.Equal(t, 123, result)
}

func TestQueueTail2(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)
	q.EnqTail(456)

	result := q.DeqTail()
	assert.Equal(t, 456, result)

	result = q.DeqTail()
	assert.Equal(t, 123, result)
}

func TestNextPointersFromEnqTail(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)
	q.EnqTail(456)
	q.EnqTail(789)

	ptr := q.head
	actual := ptr.Load()
	assert.Equal(t, 123, actual.val)

	ptr = actual.next
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 456)

	ptr = actual.next
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 789)
}

func TestPrevPointersFromEnqTail(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)
	q.EnqTail(456)
	q.EnqTail(789)

	ptr := q.tail
	actual := ptr.Load()
	assert.Equal(t, 789, actual.val)

	ptr = actual.prev
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 456)

	ptr = actual.prev
	actual = ptr.Load()
	assert.Equal(t, ptr.Load().val, 123)
}

func TestEnqHeadDeqTail1(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)
	result := q.DeqTail()
	assert.Equal(t, 123, result)
}

func TestEnqHeadDeqTail2(t *testing.T) {
	q := Queue[int]{}

	q.EnqHead(123)
	q.EnqHead(456)

	result := q.DeqTail()
	assert.Equal(t, 123, result)

	result = q.DeqTail()
	assert.Equal(t, 456, result)
}

func TestEnqTailDeqHead1(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)
	result := q.DeqHead()
	assert.Equal(t, 123, result)
}

func TestEnqTailDeqHead2(t *testing.T) {
	q := Queue[int]{}

	q.EnqTail(123)
	q.EnqTail(456)

	result := q.DeqHead()
	assert.Equal(t, 123, result)

	result = q.DeqHead()
	assert.Equal(t, 456, result)
}

func TestDeqHeadConcurrent(t *testing.T) {
	const queueSize = 1000

	q := Queue[int]{}

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < queueSize; i++ {
		q.EnqHead(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, queueSize)
	for i := 0; i < queueSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < queueSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dequed := q.DeqHead()
			m[dequed].Add(1)
		}()
	}

	wg.Wait()

	// Verify that all values from [0 to queueSize) were popped exactly once
	for i := 0; i < queueSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was popped %d times", i, count)
		}
	}
}

func TestDeqTailConcurrent(t *testing.T) {
	const queueSize = 1000

	q := Queue[int]{}

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < queueSize; i++ {
		q.EnqTail(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, queueSize)
	for i := 0; i < queueSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < queueSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dequed := q.DeqTail()
			m[dequed].Add(1)
		}()
	}

	wg.Wait()

	// Verify that all values from [0 to queueSize) were deq'ed exactly once
	for i := 0; i < queueSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was popped %d times", i, count)
		}
	}
}

func TestQueueConcurrent(t *testing.T) {
	const queueSize = 1000

	q := Queue[int]{}

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < queueSize; i++ {
		q.EnqTail(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, queueSize)
	for i := 0; i < queueSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < queueSize; i++ {
		wg.Add(1)
		go func() {
			myI := i
			defer wg.Done()

			heads := myI%2 == 0

			var dequed int
			if heads {
				dequed = q.DeqHead()
			} else {
				dequed = q.DeqTail()
			}
			m[dequed].Add(1)
		}()
	}

	wg.Wait()

	// Verify that all values from [0 to queueSize) were deq'ed exactly once
	for i := 0; i < queueSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was popped %d times", i, count)
		}
	}
}
