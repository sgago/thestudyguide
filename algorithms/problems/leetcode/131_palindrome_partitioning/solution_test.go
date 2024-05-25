package palindromepartitioning

/*
131. Palindrome Partitioning
https://leetcode.com/problems/palindrome-partitioning

Given a string s, partition s such that every substring
of the partition is a palindrome.
Return all possible palindrome partitioning of s.

Example 1:
	Input: s = "aab"
	Output: [["a","a","b"],["aa","b"]]

Example 2:
	Input: s = "a"
	Output: [["a"]]

Constraints:
- 1 <= s.length <= 16
- s contains only lowercase English letters.
*/

/*
We can DFS out all the palindromes.
I'm also seeing some potential dupes.
We'd try a + a, then a + b, and then a + b again.
But we know there's no reason to retry a + b if it
didn't work the first time. Super low amount of
s.length, 1 to 16. Indicates TC/SC of like N^2,
2^N, or N! solutions. Ah, ok, keyword here
is "substring". So, for aba, aa would not be a valid answer.

So DFS + Memo + Backtracking works here for sure.

Would there be any reason we could/should use DP?
Do we have overlap? Not...really?

*/
