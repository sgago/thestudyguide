package splitarraylargestsum

/*
410. Split Array Largest Sum
https://leetcode.com/problems/split-array-largest-sum

Given an integer array nums and an integer k,
split nums into k non-empty subarrays such that the
largest sum of any subarray is minimized.

Return the minimized largest sum of the split.
A subarray is a contiguous part of the array.

Example 1:
	Input: nums = [7,2,5,10,8], k = 2
	Output: 18
Explanation:
	There are four ways to split nums into two subarrays.
	The best way is to split it into [7,2,5] and [10,8],
	where the largest sum among the two subarrays is only 18.

Example 2:
Input: nums = [1,2,3,4,5], k = 2
Output: 9
Explanation:
	There are four ways to split nums into two subarrays.
	The best way is to split it into [1,2,3] and [4,5],
	where the largest sum among the two subarrays is only 9.

Constraints:
- 1 <= nums.length <= 1000
- 0 <= nums[i] <= 106
- 1 <= k <= min(50, nums.length)
*/

/*
So we want to minimize the sums of both k nums.
Problem is, we don't reallly know what numbers are
in ye'ol array.

One option is to sort the array (nlogn), but this doesn't
really seem to buy us anything at all. Example 1
is good cause the last two biggest numbers are actually
the smallest together.

So, we can DFS this out, a bool array of what we used.
DP looks doable. Trying to minimize something,
overlapping cause we don't know which num goes where.

   7  2  5  10  8   32 32/1
   7  9  14 24  32  16 32/2
   7  9  14 24  32  10 32/3  14 10 8? 9 15 8?
   7  9  14 24  32  8  32/4
   7  9  14 24  32  6  32/5
*/
