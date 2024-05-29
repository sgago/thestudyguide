package smalleststringwithswaps

import (
	"fmt"
	"sgago/thestudyguide/col/bitflags"
	"sgago/thestudyguide/col/dsu"
	"sgago/thestudyguide/col/queue"
	"strings"
	"testing"
)

/*
1202. Smallest String With Swaps
https://leetcode.com/problems/smallest-string-with-swaps

You are given a string s, and an array of pairs of indices in the string pairs
where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.

You can swap the characters at any pair of indices in the given pairs any number of times.
Return the lexicographically smallest string that s can be changed to after using the swaps.

Example 1:
	Input: s = "dcab", pairs = [[0,3],[1,2]]
	Output: "bacd"
	Explanation:
		Swap s[0] and s[3], s = "bcad"
		Swap s[1] and s[2], s = "bacd"

Example 2:
	Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
	Output: "abcd"
	Explanation:
		Swap s[0] and s[3], s = "bcad"
		Swap s[0] and s[2], s = "acbd"
		Swap s[1] and s[2], s = "abcd"

Example 3:
	Input: s = "cba", pairs = [[0,1],[1,2]]
	Output: "abc"
	Explanation:
		Swap s[0] and s[1], s = "bca"
		Swap s[1] and s[2], s = "bac"
		Swap s[0] and s[1], s = "abc"

Constraints:
- 1 <= s.length <= 10^5
- 0 <= pairs.length <= 10^5
- 0 <= pairs[i][0], pairs[i][1] < s.length
- s only contains lower case English letters.
*/

/*
So, we get a string s and pairs of indices of chars in the string we can swap around.
We're trying to get the string as close as possible to a to z in order (lexicographical)
as possible (like abcdefghijklmnop... and so on).

We can swap any number of times in any order to build the solution.

We need to return the smallest lexico string.

Observation: There might be a way to like merge down the pairs.
For example, given 0,3 and 1,3 we can swap any characters to 0, 1, or 3.
Maybe weird idea -> construct ranges we can run sort on, sort the chars,
reinsert the chars we can't sort on.

Like, we can convert pairs into just an array of indices we can sort on.

Another idea -> Maybe we can sum up a lexico score and try to minimize that score.

Another idea -> Maybe we can make a grid/dict of some sort to sort on.
   0 1 2 3
   d c a b
   0     3
     1 2
   0   2
   -------
   0 1 2 3

Yeah, we can loop len(pairs) times, using two slices, and build up our index array, sort
chars that appear in the array. n to build up the array of indices, nlogn to sort each chunk -> n+mlogm/c*c -> mlogm -> nlogn?
Optimization: We would only need two arrays to build the index list.

Whoa, wait, the indexes need to overlap.

Yet another idea -> Maybe we can make a graph out of this.

Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]] -> dictionary like
0[d]: 3, 2
1[c]: 2
2[a]: 1, 0
3[b]: 0

n to make the dictionary. Repeatedly DFS/BFS out the smallest lexico str? n + n*(V+E)?

Note: You can't do str := "mystring" and then str[0] = 'r' in go. Let's make a rune array and convert it and/or str builder it.

Yet back to the grid, can we build up the grid using smallest indices?
0 1 2 3
d c a b
0     0
  1 1   <- overlap
0   0
0 0 0 0 <- final, sort indices that share same number?

Like n^2 to make this, then like n*nlogn. Graph search is looking cheaper.

Opti: We can bitmask the visited col
*/

func TestSmallestStringWithSwaps(t *testing.T) {
	actual := smallestStringWithSwaps("dcab", [][]int{
		{0, 3},
		{1, 2},
		// {0, 2},
	})

	fmt.Println(actual)
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(pairs) == 0 {
		return s
	}

	if len(s) <= 1 {
		return s
	}

	sb := strings.Builder{}

	// n pairs
	graph := make(map[int][]int, len(s))
	for _, pair := range pairs {
		if len(pair) != 2 {
			panic("length of all index pairs must be exactly two elements")
		}

		if pair[0] < 0 || pair[0] > len(s) || pair[1] < 0 || pair[1] > len(s) {
			panic("pair index is not in range")
		}

		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	// m*(|V| + |E|) dfs
	done := bitflags.New(len(s))
	for i := 0; i < len(s); i++ {
		smallest := getSmallest(s, i, graph, done)
		done.Set(smallest, true)
		sb.WriteByte(s[smallest])
	}

	// Final is like n+m(V+E) ~> O(n(V+E))
	return sb.String()
}

func getSmallest(s string, key int, graph map[int][]int, done bitflags.Flags) int {
	neighbors := graph[key]

	if len(neighbors) == 0 {
		return key
	}

	smallest := -1

	visited := bitflags.New(len(graph))
	visited.Set(key, true)

	q := queue.New[int](len(graph))
	q.EnqTail(key)

	for !q.Empty() {
		curr := q.DeqHead()

		for _, idx := range graph[curr] {
			if !visited.Get(idx) {
				if !done.Get(idx) {
					if smallest == -1 || s[idx] < s[smallest] {
						smallest = idx
					}
				}

				q.EnqTail(idx)
				visited.Set(idx, true)
			}
		}
	}

	return smallest
}

/*
Ah, ok, the optimal solution uses a DSU, which approaches a n*logn + m TC.
*/

func TestOptimalSmallestStringWithSwaps(t *testing.T) {
	actual := optimalSmallestStringWithSwaps("dcab", [][]int{
		{0, 3},
		{1, 2},
		// {0, 2},
	})

	fmt.Println(actual)
}

func optimalSmallestStringWithSwaps(s string, pairs [][]int) string {
	dsu := dsu.New[int](len(s))

	for _, pair := range pairs {
		dsu.Add(pair[0], pair[1])
	}

	for _, pair := range pairs {
		dsu.Union(pair[0], pair[1])
	}

	set, _ := dsu.Set(0)

	// TODO: Get each set, sort them, sort the string indexes the same way,
	// build the string up.

	fmt.Println(set)

	return ""
}
