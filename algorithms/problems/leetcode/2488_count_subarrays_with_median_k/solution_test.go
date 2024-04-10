package countsubarrayswithmediank2488

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"

	"github.com/stretchr/testify/assert"
)

/*
2488. Count Subarrays With Median K
You are given an array nums of size n consisting of
distinct integers from 1 to n and a positive integer k.

Return the number of non-empty subarrays in nums that have a median equal to k.

Note:
The median of an array is the middle element after sorting the array in ascending order.
If the array is of even length, the median is the left middle element.
For example, the median of [2,3,1,4] is 2, and the median of [8,4,3,5,1] is 4.
A subarray is a contiguous part of an array.


Example 1:
	Input: nums = [3,2,1,4,5], k = 4
	Output: 3
	Explanation: The subarrays that have a median equal to 4 are: [4], [4,5] and [1,4,5].

Example 2:
	Input: nums = [2,3,1], k = 3
	Output: 1
	Explanation: [3] is the only subarray that has a median equal to 3.


Constraints:
	n == nums.length
	1 <= n <= 105
	1 <= nums[i], k <= n
	The integers in nums are distinct.

=====================================================================================

So, we are given a single array, and we need to find many subarrays that have a median k.
The array has distinct nums from 1 to n, and k is within 1 to n too.
For an even numbered array, the median is the left middle element.

We need to find all subarrays with median k. Note there's a very small n, suggesting N^2 or N! solution via combinatorial/backtracking/dp.

What are our options, here? A median is the middle (or left middle) element.
This means we need sorted knowledge of the numbers. There's no indication that the input is sorted.
We can sort for O(NlogN) time and then run a combinatorial DFS. We could also use heaps. For example, we could
dump numbers less than k into small heap and larger elements into large heap. I bet there's a nice math solution to this
somehow, without all the sort/heap/dfs noise.

To build out our sub arrays, we can start with k, then add right elements, then left elements such that k is middle or left of middle.

1 2 3 4 5, k=3
3
3 4
3 5
2 3 4
2 3 5
1 3 4
1 3 5
2 3 4 5
1 3 4 5
1 2 3 4 5

Ohhhh, I don't think they want me to solve it this way at all.
I think they want me to keep the array unsorted, and find subarrays within the unsorted collection.
So, no sorting permitted at all lol.
That's why my math-magic won't work here. Ok, this is very different than I first thought.
Note to self, possible math-magic means they want something else cause this is supposed to be a programming problem.

Cool, so, step one is finding our k plus numbers smaller, plus numbers larger.
We can two pointer (fast+slow) to k to capture the total subarray being considered. Like

      s  k
3, 2, 1, 4, 5
            f

Once we have 1, 4, 5. Math-magic can probably occur. Note, we should store the k index along with fast and slow pointers.

k
k Rn
Lk-1 k Rn

The math magic is going to be tricky here. We could DFS out all subarrays once we have our total subarray.
I would do a DFS combinatorial search here during an interview, 100%. Struggling with the math-magic would take me
a hot second and am unlikely to complete in an allotted chunk of time.
1 +
k R! +
L! k R!


1 2 3 4 5, k=3

3

3 4
3 5

2 3 4
2 3 5

1 3 4
1 3 5

2 3 4 5
1 3 4 5

1 2 3 4 5

It's not a simple the aggregate sum of left * right (SL * SR). With that we get (3+2+1)*(2+1)=18.
The right side is definitely (2+1), it can be any two + remaining one.
Left side needs some subtraction. It's like (3+2)*(2)+1 == 10.
(Snl-1)*(Nr)+1? Does this work? Let's move median to 4 and try again.

1 2 3 4 5, k=4

4

4 5

1 4 5
2 4 5
3 4 5

(3+2)*1+1 = 6, but it's actually 5...

Ohh oops, no +1. My b. Reminder, mental math is always dangerous lol.

Just (Snl-1)*(Nr) then? Move median to 2 and try again.

1 2 3 4 5, k=2

2

2 3
2 4
2 5

1 2 3
1 2 4
1 2 5

1 2 3 4
1 2 3 5
1 2 4 5


Two options on the left + 3 options on the right.
Left is nothing or just 1.
Right is 3, 4, 5, 3 4, 3 5, 4 5. 6 options. 3 + 2 + 1. Also nothing, but's let's ignore that for a moment.
Left is Nothing or 1. But you can't swap whatever you feel like. You don't get
1, 2, 1 2. It's just 2 and 1 2. Let's ignore k then.

+1 (just k by itself)
+3 (k + all right numbers)
+3 (1 k + all right numbers)
+3 (1 k + unique combos of 3 4 and 5)

There is a comb/perm formula we can make, I'm pretty sure.
*/

func TestAppend(t *testing.T) {
	slice := []int{1, 2, 3}
	_ = append(slice[1:1], 999)
	// Can you guess what the slice elements are now?
	fmt.Println("It's", slice)
}

type state struct {
	kIdx int
	lIdx int
	rIdx int
	nums []int
}

func (s *state) soln() bool {
	total := len(s.nums)

	if total%2 == 0 {
		return total/2-1 == s.kIdx
	}

	return total/2 == s.kIdx
}

func TestCountSubarraysWithMedianK(t *testing.T) {
	actual := countSubarraysWithMedianK([]int{3, 2, 1, 4, 5}, 4)
	fmt.Println(actual)
	assert.Equal(t, 3, actual)
}

func TestCountSubarraysWithMedianK_1to5_Kis3(t *testing.T) {
	actual := countSubarraysWithMedianK([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(actual)
	assert.Equal(t, 10, actual)
}

func countSubarraysWithMedianK(nums []int, k int) int {
	sub, idx := getSubarray(nums, k)

	count := 0

	left := sub[:idx]
	right := sub[idx+1:]

	fmt.Println("left:", left, "k:", k, "right:", right)

	memo := make(map[string]bool)

	init := state{
		kIdx: 0,
		rIdx: 0,
		lIdx: len(left) - 1,
		nums: []int{k},
	}

	q := queue.New(10, init)

	for !q.Empty() {
		curr := q.DeqHead()

		fmt.Println("curr:", curr)

		if curr.soln() {
			count++
		}

		for i := curr.lIdx; i >= 0; i-- {
			if left[i] == curr.nums[0] {
				continue
			}

			next := state{
				kIdx: curr.kIdx + 1,
				lIdx: curr.lIdx - 1,
				rIdx: curr.rIdx,
				nums: append([]int{left[i]}, curr.nums...),
			}

			key := fmt.Sprint(next.nums)

			if _, ok := memo[key]; ok {
				continue
			}

			memo[key] = true

			fmt.Println("  left:", next)

			q.EnqTail(next)
		}

		for i := curr.rIdx; i < len(right); i++ {
			if right[i] == curr.nums[len(curr.nums)-1] {
				continue
			}

			next := state{
				kIdx: curr.kIdx,
				lIdx: curr.lIdx,
				rIdx: curr.rIdx + 1,
				nums: append(curr.nums, right[i]),
			}

			key := fmt.Sprint(next.nums)

			if _, ok := memo[key]; ok {
				continue
			}

			memo[key] = true

			fmt.Println("  rght:", next)

			q.EnqTail(next)
		}
	}

	return count
}

func TestGetSubarry(t *testing.T) {
	sub, idx := getSubarray([]int{3, 2, 1, 4, 5}, 4)
	fmt.Println("sub:", sub, "| k idx:", idx)
	assert.Equal(t, 1, idx)
}

func TestGetSubarry_OneElement(t *testing.T) {
	sub, idx := getSubarray([]int{3}, 3)
	fmt.Println("sub:", sub, "| k idx:", idx)
	assert.Equal(t, 0, idx)
}

func TestGetSubarry_TwoElements3and4_Kis3(t *testing.T) {
	sub, idx := getSubarray([]int{3, 4}, 3)
	fmt.Println("sub:", sub, "| k idx:", idx)
	assert.Equal(t, 0, idx)
}

func TestGetSubarry_TwoElements3and4_Kis4(t *testing.T) {
	sub, idx := getSubarray([]int{3, 4}, 4)
	fmt.Println("sub:", sub, "| k idx:", idx)
	assert.Equal(t, 1, idx)
}

func getSubarray(nums []int, k int) ([]int, int) {
	if len(nums) == 0 {
		return nums, -1
	}

	if len(nums) == 1 {
		return nums, 0
	}

	kidx := -1
	fast, slow := 0, 0

	// 3,2,1,4,5 -> 1, 4, 5 and kidx = 1

	// TC N to be contd...
	for ; fast < len(nums); fast++ {
		if nums[fast] == k {
			kidx = fast
			break
		}

		if nums[slow] > nums[fast] {
			slow++
		}
	}

	if fast == 0 {
		fast++
	}

	// TC N contd... So just TC N on this monster
	for fast < len(nums) && fast > 0 && nums[fast-1] < nums[fast] {
		fast++
	}

	return nums[slow:fast], kidx - slow
}
