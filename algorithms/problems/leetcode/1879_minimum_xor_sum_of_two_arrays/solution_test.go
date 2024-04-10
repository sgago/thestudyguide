package minimumxorsumoftwoarrays

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"sgago/thestudyguide/col/queue"
)

/*
1879. Minimum XOR Sum of Two Arrays
https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/

You are given two integer arrays nums1 and nums2 of length n.

The XOR sum of the two integer arrays is
(nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... +
(nums1[n - 1] XOR nums2[n - 1]) (0-indexed).

For example, the XOR sum of [1,2,3] and [3,2,1]
is equal to (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 + 2 = 4.
Rearrange the elements of nums2 such that the resulting
XOR sum is minimized.

Return the XOR sum after the rearrangement.

Example 1:
	Input: nums1 = [1,2], nums2 = [2,3]
	Output: 2
	Explanation: Rearrange nums2 so that it becomes [3,2].
	The XOR sum is (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2.

Example 2:
	Input: nums1 = [1,0,3], nums2 = [5,3,4]
	Output: 8
	Explanation: Rearrange nums2 so that it becomes [5,4,3].
	The XOR sum is (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8.

Constraints:
- n == nums1.length
- n == nums2.length
- 1 <= n <= 14
- 0 <= nums1[i], nums2[i] <= 107

=================================================================

Mmk, step one, what's the XOR operator in golang? That's not one you see busted
out super often. It's ^. Quick bitwise operator review
- & is bitwise and
- | is bitwise or
- ^ is bitwise xor
- << is binary left shift
- >> is binary right shift

We want to rearrange nums2 such that the bitwise XOR sum is minimal.
Crazy. So, the goal of this is to basically sort nums2 such that
there's a maximum amount of 0's and 1's matching from the nums1[i] ^ nums2[i].
This'll keep and flip 0's and 1's respectively in our sum. Ahh, but
we're trying to minimize the sum. So we definitely want to chip away at
the most significant bits where possible. There's really a small number
of elements, making this feel like DP or similar. Still figuring that out.

So, zero has no 1's, so we want to pair that with the smallest
value possible. Likewise, the max (107) probably has the most 1's,
so we'll want to pair that up with the most amount of 1's possible.

How might we do this the brute force way?
Sorting. What does that give us?
DFS/BFS. Possible?
What's the most efficient way to count the number of one's? Is it looping?
Can we count the number of one's by bit position?
Are these overlapping subproblems? Would DP help?
Bitmasking. What does this do for us?

Another idea, what if we count the number of bits at each bit position
in nums1. Then we could try to mask them off or something? Ehh, no.

Oof, this one is hard lol. They probably want DP here, given the small n values.
Dealing with XOR is kind of annoying too.

Another idea, we can transform XOR and/or nums somehow to make this more amenable
to what's likely a DP problem. Yeah, likely DP, small n values and there's some
searching involved. Yeah, not, and, and or are functionally complete, so we
could maybe reduce this problem somehow using bitoperator magic.

Ah, right, the max number is 107. That's 0x6B in hex or
0b11010110. We can mask away bit positions 8+. We only care about
the first 7.

Again, we're try to mask off as many sig bits as possible.

nums1: 1 2 3
nums2: 3 2 1
minsum: 2 + 0 + 2 = 4

Omg, this one is tough. Ok, it's gotta be like DFS thru,
memo'ize the smallest XOR values/sums, ugh, yeah, yeah
we can basically memoize it. Top down seems tough, I don't
see a path thru immediately.

So given 1 2 3, how do we get from state 1 to 2?
1        2         3
1 is 0
2 is 3 <- the num we want
3 is 2 - min here 0 but wait...

        1 is 3
	    2 is 0 <- the num we want
	    3 is 1

			    1 is 2 <- the num we want
			    2 is 1
			    3 is 0 - min here

Yeaaaahhhh ok the bottom up is going to store the min XOR sum.
If we got to a number using a smaller value, we can use that instead.

Xor'ing zero gives the same number.
*/

func TestMinXorSumOfTwoArrays_123(t *testing.T) {
	minXorSum([]int{1, 2, 3}, []int{3, 2, 1})
}

func TestMinXorSumOfTwoArrays_1234(t *testing.T) {
	minXorSum([]int{1, 2, 3, 4}, []int{4, 3, 2, 1})
}

type state struct {
	i     int
	j     int
	sum   int
	visit int
}

func minXorSum(nums1, nums2 []int) int {
	n := len(nums1) // len(nums1) == len(nums2) from constraints
	q := queue.New[state](10)

	/*
		Lesson learned: Memoizing is simply not possible for all solutions.
		This monster's TC is n * 2^n. Why could I not memoize this though?
		We can't memoize the sum, visit value, or i,j here. A bigger sum
		at 2,0 or visit 3 could still be the minimum solution later.
		This super small slice length should probably be a giveaway too.
		Max of 14 is mega small potentially indicating a mega big TC.
		Need to think before blindly adding in a memo and wasting time.
	*/
	// useMemo := true
	// memo := make(map[int]map[int]int)

	iter := 0
	ans := math.MaxInt

	for j := 0; j < n; j++ {
		iter++

		q.EnqTail(state{
			i:   0,
			j:   j,
			sum: nums1[0] ^ nums2[j],
		})
	}

	for !q.Empty() {
		iter++

		curr := q.DeqHead()

		fmt.Println("curr i:", curr.i, "| j:", curr.j, "| sum:", curr.sum, "| vis:", curr.visit)

		// Find the next number from nums2 that we're going to add
		for j := 0; j < n-1; j++ {
			if (curr.visit & (1 << j)) == 0 {
				next := state{
					i:     curr.i + 1,
					j:     j,
					sum:   curr.sum + nums1[curr.i+1] ^ nums2[j],
					visit: curr.visit | (1 << j),
				}

				// if useMemo {
				// 	if memoSum, ok := memo[curr.i+1][j]; ok {
				// 		// Previous sum
				// 		if next.sum < memoSum {
				// 			// There's a new min sum to store in our memo
				// 			memo[curr.i+1][j] = next.sum
				// 		} else {
				// 			break // This sum is greater, so just quit early then
				// 		}
				// 	} else {
				// 		// New i,j memo value to store

				// 		if memo[curr.i+1] == nil {
				// 			// Init the memo, golang makes this hard unsure of a nicer way
				// 			memo[curr.i+1] = make(map[int]int)
				// 		}

				// 		memo[curr.i+1][j] = next.sum
				// 	}
				// }

				q.EnqTail(next)
			}
		}

		if curr.i == n-1 {
			// That was the last nums1 value to sum
			// Store the min sum in ans
			ans = min(ans, curr.sum)
		}
	}

	fmt.Println("ans:", ans)
	fmt.Println("iter:", iter)

	return ans
}

func Test107InDifferentNumberBases(t *testing.T) {
	fmt.Printf("%x", 107) // 107 == 0x6B or 0b11010110
	fmt.Println(strconv.FormatInt(int64(107), 2))
	fmt.Printf("%x", 3^3)
}
