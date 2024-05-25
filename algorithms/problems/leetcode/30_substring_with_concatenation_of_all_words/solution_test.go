package substringwithconcatenationofallwords

/*
30. Substring with Concatenation of All Words
https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

You are given a string s and an array of strings words.
All the strings of words are of the same length.

A concatenated string is a string that exactly contains
all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef",
"abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are
all concatenated strings. "acdbef" is not a concatenated
string because it is not the concatenation of any permutation
of words.

Return an array of the starting indices of all the
concatenated substrings in s. You can return the answer in
any order.

Example 1:
	Input: s = "barfoothefoobarman", words = ["foo","bar"]
	Output: [0,9]
	Explanation:
	The substring starting at 0 is "barfoo".
		It is the concatenation of ["bar","foo"] which is a permutation of words.
		The substring starting at 9 is "foobar".
		It is the concatenation of ["foo","bar"] which is a permutation of words.

Example 2:
	Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
	Output: []
	Explanation:
		There is no concatenated substring.

Example 3:
	Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
	Output: [6,9,12]
	Explanation:
		The substring starting at 6 is "foobarthe".
		It is the concatenation of ["foo","bar","the"].
		The substring starting at 9 is "barthefoo".
		It is the concatenation of ["bar","the","foo"].
		The substring starting at 12 is "thefoobar".
		It is the concatenation of ["the","foo","bar"].

Constraints:
- 1 <= s.length <= 104
- 1 <= words.length <= 5000
- 1 <= words[i].length <= 30
- s and words[i] consist of lowercase English letters.
*/

/*
Given abcdefcba, and a, b, and c, we return an array of indexes 0 and 6
because we can make abc and cba.

Mmk, so, we can DFS out all the combos.
If we have w subwords, then there's w! combos to consider.
Not super exciting. Note, we need unique combos, not dupes.

It's the match part we need to think about.
We can simply generate all combos and do a strings.Index(s, substr) bool call.
Gives us like an W!*S vibe.

The question is, can we bake in some pruning somehow?
We could memo out some values. But what to memo?

Say I memo, a b c with found indexes like 0, 1, 2. Note these are first
indicies, not all of them. There are dupes.
If any a b or c is not found in abcdefcba, then we can quit immediately.
We have to match the whole perm, not just part of it, pretty sure.
Yeah "exactly contains all strings".

Then I memo, ab ac ba bc cb ca with indexes 0, -1, 7, 1, 6.
We can prune away ac cause it doesn't exist.
Getting rid of ac removes acb only.

Is the memo actually faster? We have to do
W*N comparisons, then W(2)*N-X comparisons, and so on.

What if we actually stored all indexes in the memo? Would build like

012345678
abcdefcba

 a 0 8 <- Same, we can drop indicies > 6
 b 1 7 <- Same, we can drop indicies > 6
 c 2 6
ab 0
ac NONE
ba 7 <- We need not consider indexes that are > N-len(words) or > 6
bc 1
ca NONE
cb 6

With this improvement we get
012345678
abcdefcba
 a 0
 b 1
 c 2 6
ab 0
ac NONE
ba NONE
bc 1
ca NONE
cb 6

We'd only consider then abc, bca, cba for full matches at indexes 0, 1, and 6 only.

What else? def are show stoppers too. Any patterns that are
index(d)-len(words) or > 0, I guess, can be skipped, too. With that, we'd build our memo as

Our no-go indexes are 1, 2, 3, 4, 5, 7, 8. Skip patterns that would start or contain these.

012345678
abcdefcba
 a 0
 b NONE
 c 6
ab 0
ac NONE
ba NONE
bc NONE
ca NONE
cb 6
...
abc 0 done
cba 6 done

Will this work? How do we build our go, nogo list?
Reminder: a b c could be full words.

Init. Add indicies 7 and 8 to naughty list. No sequence may start here.

Find a at 0. Look 2 chars up. See c. c is in a b c. Good.
Stop looping at N-len(words).
Find b at 1. Look 2 chars up. See d. d is not in a b c. Bad. Add 1 and 3 to naughty list? No sequence can start or contain these.
Stop looping at N-len(words).
Find c at 2. Look 2 chars up. See e. e is not in ab c c. Bad. Add 2 and 4 to naughty list. No sequence can start or contain these.

Check soln now. Ahh ok, close. Should have kept going.
I had like 60-70% of the soln here. For some reason, I thought a DFS was still necessary.

I could use my bitwise collection, and only use two 64 bit numbers to store a good/not-good value.

*/
