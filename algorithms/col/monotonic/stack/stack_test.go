package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := New[int](2)

	lis := 0

	for _, num := range []int{0, 1, 3, 2, 4, 5, -1, 0, 3} {
		lis = max(lis, len(s.Push(num))+s.Len()-1)
	}

	fmt.Println(lis)
}
