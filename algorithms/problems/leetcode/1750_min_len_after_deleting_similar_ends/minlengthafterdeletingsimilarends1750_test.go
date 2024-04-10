package leetcode

/*
1750. Minimum Length of String After Deleting Similar Ends
https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/

Given a string s consisting only of characters 'a', 'b', and 'c'.

You are asked to apply the following algorithm on the string any number of times:
	1. Pick a non-empty prefix from the string s where all the characters in the prefix are equal.
	2. Pick a non-empty suffix from the string s where all the characters in this suffix are equal.
	3. The prefix and the suffix should not intersect at any index.
	4. The characters from the prefix and suffix must be the same.
	5. Delete both the prefix and the suffix.

Return the minimum length of s after performing the above operation any number of times (possibly zero times).

Example 1:
	Input: s = "ca"
	Output: 2
	Explanation: You can't remove any characters, so the string stays as is.

Example 2:
	Input: s = "cabaabac"
	Output: 0
	Explanation: An optimal sequence of operations is:
	- Take prefix = "c" and suffix = "c" and remove them, s = "abaaba".
	- Take prefix = "a" and suffix = "a" and remove them, s = "baab".
	- Take prefix = "b" and suffix = "b" and remove them, s = "aa".
	- Take prefix = "a" and suffix = "a" and remove them, s = "".

Example 3:
	Input: s = "aabccabba"
	Output: 3
	Explanation: An optimal sequence of operations is:
	- Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
	- Take prefix = "b" and suffix = "bb" and remove them, s = "cca".

Constraints:
	1 <= s.length <= 105
	s only consists of characters 'a', 'b', and 'c'.

================================================================================================

Mmk, this has a two-pointer moving in opposite directions vibe. Super weird problem.
Basically, if a string has starting and ending characters that are the same, we remove them.
1. ccccbbbacbbccc - Trim 4 c's on left and 3 on right, b/c both prefix and suffix have same chars.
2. bbbacbb - Trim 3 b's on left and  2 on right, b/c again both prefix and suffix have same chars.
3. ac - Nothing to do, a != c so we can't trim any more. Output 2.

Start with a left ptr at 0 and a right ptr at len(the_passed_in_string)-1. Walk them in while chars equal.
Cut a slice from left to right, watching the [left, right) inclusive/exclusive stuff, then loop around and do it again.
Return len(the_trimmed_string). Also, we need to quit if left == right, so watch that. We'll need to store which character we care about.
Test cases:
- INPUT, OUTPUT
- empty string, 0
- a, 1 - a can't be removed, left == right
- ac, 2 - a != c, cannot be removed at all
- aa, 0 - all can be removed
- cac, 1 - a can't be removed, left == right
- aaaa, 0 - all can be removed.
- aaaaa, 1 - all but the middle a can be removed, left == right
- ccac, 1 - a can't be removed
- cabaabac, 0 - all can be removed.
- aabccabba, 3 - cca different prefix/suffix, rest can be removed
- caabccabba, 10 - different prefix/suffix
*/
