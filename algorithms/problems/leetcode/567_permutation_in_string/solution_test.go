package permutationinstring

/*
567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/?envType=daily-question&envId=2024-10-05

Given two strings s1 and s2, return true if s2 contains a  permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.

Example 1:
	Input: s1 = "ab", s2 = "eidbaooo"
	Output: true
	Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
	Input: s1 = "ab", s2 = "eidboaoo"
	Output: false

Constraints:
- 1 <= s1.length, s2.length <= 104
- s1 and s2 consist of lowercase English letters.
*/

/*
  So, basically, we check to see if we can rearrange the characters in s1 to form a substring of s2.
  Looks like a two part-er, first we can create permutations of s1 and then check if any of them are substrings of s2.
  We need to create perms that contain all the characters of s1, so length needs to be the same.

  What else? Maybe we could sum the characters and do a prefix sum of sorts? Maybe add them to a map?
  Yeah, I could add them to a map, summing each character and slide along the string, removing the first character and adding the next.
  If the map is the same, we have a match.
  It's N to create the map, and M to slide along the string, so O(N+M).
  If s1 is larger than s2, quit immediately b/c there's no answer. Ok, the s1 will be the same size or smaller so nvm.

  Looks like Go has a builtin DeepEqual function that can compare maps.
  Generating permutations is like N!*N. Creating the map is N, the sliding window is N * M, and the map comparison is N * M.
  Creating the s1 and s2 map is N * 2. Then we need to slide past M characters so 2N*M. Each M, we need to compare the maps.
  So this is N^2*M, I think? That should be better than N!*N for generating permutations plus the dropped N terms.

  Oh, looks like this is the correct approach but wrong TC, apparently its O(N+M). Ah, it's because the
  map will only have 26 entries, so it's not N. The sliding window is M, so it's N+M. Cool.
*/
