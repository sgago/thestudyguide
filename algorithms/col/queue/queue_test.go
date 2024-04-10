package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int](5, 1, 2, 3)

	q.EnqHead(4, 5, 6)
	q.EnqTail(7, 8, 9)
	q.DeqHead()
}
