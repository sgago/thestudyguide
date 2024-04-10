package interval

import (
	"fmt"
	"testing"
)

func TestMergeAll(t *testing.T) {
	intervals := []Interval[int]{
		New(0, 9),
		New(5, 15),
		New(10, 19),
		New(20, 29),
		New(-1, 0),
	}

	merged := MergeAll(intervals...)
	fmt.Println(merged)
}
