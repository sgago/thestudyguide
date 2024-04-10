package backtracking

import (
	"testing"

	"sgago/thestudyguide/col/queue"
)

const maxLength = 4

var letters []string = []string{"a", "b"}

func TestBacktracking1(t *testing.T) {
	backtracking1()
}

func backtracking1() *[]string {
	q := queue.New[string](2, letters...)
	soln := queue.New[string](4)

	for !q.Empty() {
		str := q.DeqTail()

		for _, letter := range letters {
			next := str + letter

			if len(next) == maxLength {
				soln.EnqTail(next)
			} else {
				q.EnqTail(next)
			}

		}
	}

	s := soln.Slice()

	return &s
}
