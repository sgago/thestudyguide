package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTrueInSortedBooleanArray_Odd(t *testing.T) {
	bools := []bool{false, false, false, true, true}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, 3, actual)
}

func TestFirstTrueInSortedBooleanArray_Even(t *testing.T) {
	bools := []bool{false, false, true, true}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, 2, actual)
}

func TestFirstTrueInSortedBooleanArray_NoTrue(t *testing.T) {
	bools := []bool{false, false, false, false}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, -1, actual)
}

func TestFirstTrueInSortedBooleanArray_AllTrue(t *testing.T) {
	bools := []bool{true, true, true, true}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, 0, actual)
}

func TestFirstTrueInSortedBooleanArray_SingleTrue(t *testing.T) {
	bools := []bool{true}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, 0, actual)
}

func TestFirstTrueInSortedBooleanArray_SingleFalse(t *testing.T) {
	bools := []bool{false}

	actual := firstTrueInSortedBooleanArray(bools)
	fmt.Println(actual)

	assert.Equal(t, -1, actual)
}

func firstTrueInSortedBooleanArray(bools []bool) int {
	ans := -1
	l, r := 0, len(bools)-1

	for l <= r {
		m := (r + l) >> 1

		if !bools[m] {
			l = m + 1
		} else {
			ans = m
			r = m - 1
		}
	}

	return ans
}
