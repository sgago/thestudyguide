package heap

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/heap"
	"sgago/thestudyguide/col/vec"
)

/*
Given a list of points on a 2D plane. Find k closest points to origin (0, 0).
	Input: [(1, 1), (2, 2), (3, 3)], 1
	Output: [(1, 1)]
*/

func TestKClosestPoints(t *testing.T) {
	// Optimization: wrap Vector2 in other struct and compute the distance only once
	vectors := []vec.Vec2{
		vec.New2D(3, 3),
		vec.New2D(2, 2),
		vec.New2D(1, 1),
		vec.New2D(4, 4),
		vec.New2D(5, 5),
	}

	less := func(v1, v2 vec.Vec2) bool {
		return v1.DistOrigin() < v2.DistOrigin()
	}

	k := 5
	h := heap.NewFunc[vec.Vec2](3, less, vectors...)

	for i := 0; i < k; i++ {
		v := h.Pop()
		fmt.Printf("%v:%f\n", v, v.DistOrigin())
	}
}
