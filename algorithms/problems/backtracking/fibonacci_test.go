package backtracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Memoization
	https://algo.monster/problems/memoization_intro

	Fibonacci number via DFS + backtracking + memo.
*/

func TestFibWithoutMemo(t *testing.T) {
	fib := fibWithoutMemo(5)
	fmt.Println("fib5 wihtout memo", fib)
	assert.Equal(t, 5, fib)
}

/*
	 Fib numbers without memo, look how many
	 duplicate n we end up using!
		n: 5
		n: 4
		n: 3
		n: 2
		n: 1
		n: 0
		n: 1
		n: 2
		n: 1
		n: 0
		n: 3
		n: 2
		n: 1
		n: 0
		n: 1
*/
func fibWithoutMemo(n int) int {
	fmt.Println("n:", n)

	if n <= 1 {
		return n
	}

	return fibWithoutMemo(n-1) + fibWithoutMemo(n-2)
}

/*
And with memo, we have:

	n: 5
	n: 4
	n: 3
	n: 2
	1/0 1
	1/0 0
	1/0 1
	cached hit: 1
	cached hit: 2
	fib5 with memo 5

	Buuuut, we do use more mem to make this happen!
	Tradeoffs, tradeoffs, tradeoffs.
*/
func TestFibWithMemo(t *testing.T) {
	memo := map[int]int{}
	fib := fibWithMemo(5, &memo)

	fmt.Println("fib5 with memo", fib)
	assert.Equal(t, 5, fib)
}

func fibWithMemo(n int, memo *map[int]int) int {
	if n <= 1 {
		fmt.Println("1/0", n)
		return n
	}

	// See if we've got a cached hit in the memo
	// from a previous time?
	if hit, ok := (*memo)[n]; ok {
		fmt.Println("cached hit:", hit)
		return hit
	}

	fmt.Println("n:", n)
	(*memo)[n] = fibWithMemo(n-1, memo) + fibWithMemo(n-2, memo)
	return (*memo)[n]
}

/*
	Whoa, I think I see DP vs DFS + memo + backtracking now...
	Let's try to do LIS with DFS + memo'ing again...
	Again, let's just memo the soln to the subproblems
	which sum up to the final soln. So, we'll memo the current
	sequence len.
*/

func TestLisDfsWithoutMemo(t *testing.T) {
	nums := []int{1, 2, 4, 3, 2, 1, 5, 6, 7}
	actual := lisDfsWithoutMemo(nums, len(nums)-1)
	fmt.Println(memolessCall, actual)
}

var memolessCall int

func lisDfsWithoutMemo(nums []int, idx int) int {
	memolessCall++

	if len(nums) <= 1 {
		return 1
	}

	lis := 1

	for i := idx - 1; i >= 0; i-- {
		if nums[idx] > nums[i] {
			curr := lisDfsWithoutMemo(nums, i) + 1
			lis = max(lis, curr)
		}
	}

	return lis
}

func TestLisDfsWithMemo(t *testing.T) {
	nums := []int{1, 2, 4, 3, 2, 1, 5, 6, 7}
	memo := map[int]int{}
	actual := lisDfsWithMemo(nums, len(nums)-1, &memo)

	fmt.Println(memoCall, actual)
}

var memoCall int

func lisDfsWithMemo(nums []int, idx int, memo *map[int]int) int {
	memoCall++

	if len(nums) <= 1 {
		return 1
	}

	if hit, ok := (*memo)[idx]; ok {
		return hit
	}

	lis := 1

	for i := idx - 1; i >= 0; i-- {
		if nums[idx] > nums[i] {
			curr := lisDfsWithMemo(nums, i, memo) + 1
			lis = max(lis, curr)
		}
	}

	(*memo)[idx] = lis
	return lis
}
