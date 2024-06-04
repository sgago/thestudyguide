package counttripletsthatcanformtwoarrysofequalxor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1442. Count Triplets That Can Form Two Arrays of Equal XOR
https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor

Given an array of integers arr.

We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).

Let's define a and b as follows:
- a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
- b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]

Note that ^ denotes the bitwise-xor operation.

Return the number of triplets (i, j and k) Where a == b.

Example 1:
	Input: arr = [2,3,1,6,7]
	Output: 4
	Explanation:
		The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)

Example 2:
	Input: arr = [1,1,1,1,1]
	Output: 10

Constraints:
- 1 <= arr.length <= 300
- 1 <= arr[i] <= 108
*/

/*
Restating the index constraints cause they're hard to read:
- 0 <= i < j
- i < j <=k
- j <= k < arr.length
- k < arr.length

Put another way
- i is in range [0,j)
- j is in range (i,k]
- k is in range [j,len)

Ok, and we're trying to pick these indices in arr
such that a == b when we xor all the elements.
Rather, we want a count of all the indices where a==b.
XOR rules:
0 ^ 0 = 0
0 ^ 1 = 1
1 ^ 0 = 1
1 ^ 1 = 0
01 ^ 11 = 10

Brute force is lots of looping, retrying. Seems like more work
to make the brute solution than something else with two pointer or dp lol.

Can we un-XOR somehow? Like, can I prefix sum this somehow?
2 ^ 3 ^ 4
010 ^ 011 ^ 100 = 001 ^ 100 = 101
Can we reverse out 2 from 101?
XORfix = 010, 001, 101
101^010 = 111, oooo that's 3 ^ 4. Yep, I think this works.

This is great, but our a/b slices can be anywhere from i,j to j,k.
We sort of end up with 2 loops, outer has an i range, then
inner holds a j range, and then the final one moves k.
That's 3 loops for like a n^3 soln.

Lots of repeat calcs. We can memo the ranges.
i   j k
2 3 1 6 7

Weird, can we DP/DFS/2D out solns more efficiently somehow?

  0 1 2 3 4
  2 3 1 6 7
2 2 1 0 6 1
3   3 2 5 2
1     1 7 0
6       6 1
7         7

0 to i == 0 to max-1
    0 1 2 3
    2 3 1 6
0 2 2 1 0 6
1 3   3 2 5
2 1     1 7
3 6       6

1 to j == 1 to max
    1 2 3 4
    3 1 6 7
1 3 3 2 4 3
2 1   1 7 6
3 6     6 1
4 7       7

So,
	i == 0, nums[i] == 2
	j == 1, any numbers in 1 == 2? Yes, at 1,2 == 3 ^ 1. Soln at 0, 1, 2.
Eh, maybe overkill with all these collections.
The xorfix 2 1 0 6 1 should be all we need to find any range we want
with a xorfix[right]^xorfix[left-1]. That's a relatively cheap op.

So, maybe it's not n^3? Probably wrong on the soln tc.
Think it's n^2. sc is like n.

So, I'm def not going to name my go slice "arr".
It's like naming a number str or something.
*/

func TestCountTriplets_23167(t *testing.T) {
	actual := countTriplets([]int{2, 3, 1, 6, 7})

	fmt.Println(actual)

	assert.Equal(t, 4, actual)
}

func countTriplets(nums []int) int {
	ans := 0
	n := len(nums)
	xorfix := make([]int, n)
	xorfix[0] = nums[0]

	for i := 1; i < n; i++ {
		xorfix[i] = nums[i] ^ xorfix[i-1]
	}

	for left := 0; left < n-1; left++ {
		mid := left + 1
		a := xorfix[left]

		for right := mid; right < n; right++ {
			// Gives xor range i to j
			b := xorfix[mid-1] ^ xorfix[right]

			if a == b {
				ans++
			}
		}

		// TODO: We need to check mid to right, then right to mid
		// This is only half the soln.
	}

	return ans
}
