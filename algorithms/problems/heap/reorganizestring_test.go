package heap

import (
	"fmt"
	"sgago/thestudyguide/col/heap"
	"testing"
)

type Pair struct {
	R rune
	N int
}

func TestReorganizeString(t *testing.T) {
	s := "aaab"
	freq := make(map[rune]int, len(s))

	for _, r := range s {
		freq[r]++
	}

	pairs := make([]Pair, 0, len(freq))
	for r, n := range freq {
		pairs = append(pairs, Pair{R: r, N: n})
	}

	less := func(p1, p2 Pair) bool {
		return p1.N > p2.N
	}

	h := heap.NewFunc(0, less, pairs...)

	for _, _ = range pairs {
		fmt.Println(h.Pop())
	}
}
