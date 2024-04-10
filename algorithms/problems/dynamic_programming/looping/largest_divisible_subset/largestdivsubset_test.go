package dynamic

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Largest Divisible Subset (LDS)
	https://algo.monster/problems/largest_divisible_subset

	Mmk, look out cause crunching the recurrence relation is going
	to be PG-13 for sure. This ain't for small children.

	Let's use 4, 3, 2, 1 as our array. It's small enough my brain can
	handle it and has some good cases like 3/4 and what not.

	For bottom up, we're going to find subsets where all the numbers are divisble
	by each other. So, [1, 2, 4] and [1, 3] and [1, 2] and [1] and [3] are
	all valid subsets.

	Remember, the goal of dynamic programming is to break problems down
	into subproblems and store the results of those subproblems into
	some dp[i] or memo[i] to avoid recalculating. Ideally, we should
	only calculate each subproblem once.

	Per usual, we should identify what we're storing into our memo.
	We're looking for the longest sequence, so we'll store the max sequence
	length into the memo. There's really nothing else of note to store that
	helps us out, indicies don't make sense or anything like that.

	Next up, the smallest subproblem is going to
	be sets of length 1. 1%1=0, 2%2=0, etc. So, by default, the min value
	of memo[i] will be 1.

	So, let's look at just element 0, the number 4. 4%4 is 0, so
	memo[0] = 1. Riveting.

	How the !@#$ do we get the result of memo[1] for the next number
	number 3?
	Both 3%4 = 3 and 4%3 = 1, so 4 ain't divisible by 3 or vice versa.
	There's no more numbers to check before 4, so our max sequence length
	is still 1.

	Now, how do we get memo[2]? This one is more exciting.
	Just like before, 2%3 and 3%2 are both > 0, so 3 and 2 aren't divisible or
	vice versa. We pass over 3. However, we now get to check 2%4 and 4%2. Ooo, 4%2 == 0.
	Exciting. So, that means 4 is divisible by 2. We get a length two
	subset. Hot damn.

	So, what'd we do to get here? We do
	memo[i] = max(memo[i-1], memo[i-2]... memo[0])+1
	but ONLY take the max of numbers where
	nums[i]%nums[i-1]==0 or nums[i-1]%nums[i]==0,
	nums[i]%nums[i-2]==0 or nums[i-2]%nums[i]==0,
	...
	nums[i]%nums[0]==0 or nums[0]%nums[i]==0.

	Can we actually say that? Does this work for weird subsets.
	Say we had 6, 3, 9, 1.
	6|1
	3|2
	9|... not divisble by 6... it's supposed to be 2, but the above will make it 3!
	1|3 cause 1 is identitiy and divides into everything without a remainder.

	If we sorted to 1, 3, 6, 9, then this'll work out then.
	6 and 9 would both have lengths of 3.
	It also nicely simplifies the
	recurence relation to:
	memo[i] = max(memo[i-1], memo[i-2]... memo[0])+1
	but ONLY the max of numbers where
	nums[i]%nums[i-1]==0,
	nums[i]%nums[i-2]==0,
	...
	nums[i]%nums[0]==0.

	Thinking again about our memo, pretty sure we don't want to memoize
	indicies or anything like that. Each pair is going to be unique cause
	it's a set.

	Lessons learned:
	- Keep sorting or changing the input array in mind. Especially for potential N^2 or 2^N solutions.
	- I liked how this flowed, it was slow but steady. Probably too slow for interview but whatever.
	- I liked authoring the recurrence relation in sentence form. I can
	  clearly understand what I need to do from it.
	- The small list of nums being ~1000 seems to suggest DP and >N^2 result.
*/

func TestLds123(t *testing.T) {
	nums := []int{1, 2, 3}
	actual := largestDivisibleSubset(nums)
	assert.Equal(t, 2, actual)
}

func TestLds1248(t *testing.T) {
	nums := []int{1, 2, 4, 8}
	actual := largestDivisibleSubset(nums)
	assert.Equal(t, 4, actual)
}

func largestDivisibleSubset(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slices.Sort(nums)

	result := 1
	memo := make([]int, len(nums))
	memo[0] = 1

	// N
	for i := 1; i < len(nums); i++ {
		curr := 1

		// Inner loop dependent on outer, so N
		// N*N = N^2
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				curr = max(curr, memo[j]+1)
			}
		}

		memo[i] = curr
		result = max(curr, result)
	}

	return result
}
