package coingame

import (
	"fmt"
	"sgago/thestudyguide/col/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinGame(t *testing.T) {
	actual := coinGame([]int{4, 4, 9, 4})

	fmt.Println(actual)

	assert.True(t, actual)
}

func coinGame(coins []int) bool {
	n := len(coins)

	dp := grid.New[int](n, n)

	for i, c := range coins {
		dp.Set(i, i, c)
	}

	fmt.Println(dp)

	for z := 1; z < n; z++ {
		r, c := 0, z

		for c < n {
			left := dp.Get(r, c-1)
			down := dp.Get(r+1, c)

			dp.Set(r, c, max(left, down))
			fmt.Println(dp)

			r++
			c++
		}
	}

	panic("")
}
