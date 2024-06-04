package singlenumberiii

/*
260. Single Number III
https://leetcode.com/problems/single-number-iii

Given an integer array nums, in which exactly two elements
appear only once and all the other elements appear exactly twice.
Find the two elements that appear only once.
You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity
and uses only constant extra space.

Example 1:
	Input: nums = [1,2,1,3,2,5]
	Output: [3,5]
	Explanation:  [5, 3] is also a valid answer.

Example 2:
	Input: nums = [-1,0]
	Output: [-1,0]

Example 3:
	Input: nums = [0,1]
	Output: [1,0]

Constraints:
- 2 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 2^31 - 1
- Each integer in nums will appear twice, only two integers will appear once.
*/

/*
So brute forces are like two loops (N^2) and or sorting NLogN + N search.
But O(N) though? Tricky. This has some cycle sort vibes to it, maybe?

Ah, man, yeah, this is an XOR bit manipulation. Of course, leetcode loves their
XOR problems. So, XOR has the following properties:
a ^ a = 0
a ^ 0 = a

Two integers will only appear once. These numbers are different, so
we can begin to form a solution from this.
*/
