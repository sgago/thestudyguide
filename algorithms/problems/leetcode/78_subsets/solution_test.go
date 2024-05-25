package subsets

import (
	"fmt"
	"sgago/thestudyguide/col/queue"
	"testing"
)

/*
78. Subsets
Given an integer array nums of unique elements, return all possible
subsets (the power set).

The solution set must not contain duplicate subsets.
Return the solution in any order.

Example 1:
	Input: nums = [1,2,3]
	Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
	Input: nums = [0]
	Output: [[],[0]]

Constraints:
- 1 <= nums.length <= 10
- -10 <= nums[i] <= 10
- All the numbers of nums are unique.
*/

/*
So, we can DFS out all combinations.
This problem has unique concerns - we cannot output both 1,2 and 2,1 together.
Either 1,2 or 2,1 but not both.
We need to track whether a number has been used or not.
We can track this using bit manipulation or just a dict of bools.
We can loop or recurse.
Nums is only up to 10 long, bitflag struct can be uint16
*/

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	sets := subsetsLooping(nums)
	fmt.Println(sets)
}

type state struct {
	nums []int
	idx  int
}

func subsetsLooping(nums []int) [][]int {
	ans := make([][]int, 0, 1<<len(nums))

	q := queue.New[state](len(nums))

	q.EnqTail(state{
		nums: []int{},
		idx:  -1,
	})

	for !q.Empty() {
		curr := q.DeqHead()

		ans = append(ans, curr.nums)

		for i := curr.idx + 1; i < len(nums); i++ {
			next := state{
				nums: curr.nums,
				idx:  i,
			}

			next.nums = append(next.nums, nums[i])

			q.EnqTail(next)
		}
	}

	return ans
}

func TestSubsetsRecursion(t *testing.T) {
	nums := []int{1, 2, 3}

	ans := make([][]int, 0, 1<<len(nums))

	subsetsRecursion(nums, []int{}, 0, &ans)

	fmt.Println(ans)
}

func subsetsRecursion(nums []int, curr []int, idx int, ans *[][]int) {
	*ans = append(*ans, curr)

	for i := idx; i < len(nums); i++ {
		next := append(curr, nums[i])

		subsetsRecursion(nums, next, i+1, ans)
	}
}
