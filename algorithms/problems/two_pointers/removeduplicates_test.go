package twopointers

import "testing"

/*
      s
	0 0 0 1 1 1 2 2
	        f
*/

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 0, 1, 1, 1, 2, 2}

	removeDuplicates(&arr)
}

func removeDuplicates(arr *[]int) int {
	a := *arr

	if len(a) <= 1 {
		return len(a)
	}

	slow, fast := 0, 1

	for ; fast < len(a); fast++ {
		if a[slow] != a[fast] {
			slow++
			a[slow] = a[fast]
		}
	}

	*arr = (*arr)[:slow+1]

	return slow + 1
}
