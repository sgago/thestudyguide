package minimumnumberofswapstomakethestringbalanced

/*
1963. Minimum Number of Swaps to Make the String Balanced
https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced

You are given a 0-indexed string s of even length n.
The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.
A string is called balanced if and only if:
- It is the empty string, or
- It can be written as AB, where both A and B are balanced strings, or
- It can be written as [C], where C is a balanced string.

You may swap the brackets at any two indices any number of times.
Return the minimum number of swaps to make s balanced.

Example 1:
	Input: s = "][]["
	Output: 1
	Explanation: You can make the string balanced by swapping index 0 with index 3.
	The resulting string is "[[]]".

Example 2:
	Input: s = "]]][[["
	Output: 2
	Explanation: You can do the following to make the string balanced:
	- Swap index 0 with index 4. s = "[]][][".
	- Swap index 1 with index 5. s = "[[][]]".
	The resulting string is "[[][]]".

Example 3:
	Input: s = "[]"
	Output: 0
	Explanation: The string is already balanced.

Constraints:
- n == s.length
- 2 <= n <= 106
- n is even.
- s[i] is either '[' or ']'.
- The number of opening brackets '[' equals n / 2, and the number of
  closing brackets ']' equals n / 2.
*/

/*
Mmk, so this is a parens game where we need to count the minimum number of swaps
to make it "balanced". Balanced means that the parens are in the correct order.
We can swap any two parens at any time. The string is guaranteed to have
opening and closing brackets in equal number. We don't need to worry about loners.
An AB balanced string is like [][] I guess? And a [C] balanced string is like [[][]]?
I wish the showed both an AB and [C] example in the prompt, it doesn't feel clear.

I don't see an immediate solution to this. What are our options? We're doing a count of swaps.
Each ] without a [ before it needs to be swapped, 100%.
Could we start in the middle and work our way out? No, doesn't look like it.
We could start on the left and right, and work our way in.
We can find the first flip-flop pair and swap them, then find the next flip-flop pair and swap those.
I don't think that works either. We'd end up with 3 moves.

Hmm, didn't see the solution to this one immediately. The solution uses greedy algorithms.
Time to study this upside down and sideways: https://www.youtube.com/watch?v=-WTslqPbj7I&t=6s
*/
