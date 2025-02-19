package constructsmallestnumberfromdistring

import "testing"

/*
2375. Construct Smallest Number From DI String
https://leetcode.com/problems/construct-smallest-number-from-di-string/

You are given a 0-indexed string pattern of length n consisting of the characters
'I' meaning increasing and 'D' meaning decreasing.

A 0-indexed string num of length n + 1 is created using the following conditions:
- num consists of the digits '1' to '9', where each digit is used at most once.
- If pattern[i] == 'I', then num[i] < num[i + 1].
- If pattern[i] == 'D', then num[i] > num[i + 1].

Return the lexicographically smallest possible string num that meets the conditions.

Example 1:
	Input: pattern = "IIIDIDDD"
	Output: "123549876"
	Explanation:
		At indices 0, 1, 2, and 4 we must have that num[i] < num[i+1].
		At indices 3, 5, 6, and 7 we must have that num[i] > num[i+1].
		Some possible values of num are "245639871", "135749862", and "123849765".
		It can be proven that "123549876" is the smallest possible num that meets the conditions.
		Note that "123414321" is not possible because the digit '1' is used more than once.

Example 2:
	Input: pattern = "DDD"
	Output: "4321"
	Explanation:
		Some possible values of num are "9876", "7321", and "8742".
		It can be proven that "4321" is the smallest possible num that meets the conditions.

Constraints:
- 1 <= pattern.length <= 8
- pattern consists of only the letters 'I' and 'D'.

================================================================================

Ok, so we get a string of I's and D's and we need to construct the smallest number possible.
We can only use each digit once and digits can only be 1 to 9.

We cannot decrease past 1 and we cannot increase past 9. A stack is interesting here.
We can push the digits to the stack and when we encounter a 'D' we can pop the stack until we find an 'I'.
This way we can keep the stack sorted in increasing order and when we encounter a 'D' we can pop the stack.

1 2 3 5 4 9 8 7 6
I I I D I D D D

3 2 1 5 4 6 7 8 9
D D D I D I I I

Our numbers are either increasing or decreasing.

One thing we could do is keep a slice of the numbers. We point to the the first/last element of the slice.
Actually, there's not that many patterns. We can just try all of them and see which one is the smallest.

We can also DFS out the answer too. We can try all the numbers and see which one is the smallest.
We'll need to keep track of which numbers we've used.
*/

func TestSmallestNumberFromDiString(t *testing.T) {
	SmallestNumberFromDiString("IIIDIDDD")
}

func SmallestNumberFromDiString(di string) {

}
