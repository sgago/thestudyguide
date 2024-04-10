package backtracking

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/queue"
)

/*
Generate All Valid Parentheses
https://algo.monster/problems/generate_parentheses

Given an integer n, generate all strings with n matching parentheses. "matching" parentheses mean
there is equal number of opening and closing parentheses.
each opening parenthesis has matching closing parentheses.
For example, () is a valid string but )( is not a valid string because ) has no matching parenthesis
before it and ( has no matching parenthesis after it.

Input
n: number of matching parentheses
Output
all valid strings with n matching parentheses

Examples

Example 1:
Input:
n = 2
Output: (()) ()()

Explanation:
There are two ways to create a string with 2 matching parentheses.

Example 2:
Input:
n = 3
Output: ((())) (()()) (())() ()(()) ()()()

Explanation:
There are 5 ways to create a string with 3 matching parentheses.

Notes:
Input: n=1
Output: ()

( -> No matching, note strlen is odd
(() -> Invalid, no matching, odd again
()) -> Invalid, no matching, odd again
() -> valid, even
()() -> valid, even
(()()) -> valid, even
The left most side is always a (, the right side is always a )
Final length of string is always 2*n. Number of ( and ) will be n.

Wow, this is fun. So, where to start? We want to generate all combinations, so this is a dfs for sure.

So, for n=1...

Left = 2, right = 2

0|   (    // Always root, sub 1 from left, left = 1, right = 2
1| (   )  // Left side, we can decr 1 from left, left = 0, right = 2 == ((. Right side: we can decr 1 from right; left = 1, right = 1 == ()
2| )   (  // Left side is out of (, can only use ) now. Decr 1 from right, left = 0, right = 1 == ((). Right side, we've got both left and right. Always left first, so decr left. left = 0, right = 1 == ()
3| )   )
*/

type state struct {
	str   string
	left  int
	right int
}

func TestGenerateAllValidParentheses(t *testing.T) {
	depth := 4 // Our "n" value

	soln := generateAllValidParentheses(depth)

	for _, str := range soln {
		fmt.Println(str)
	}
}

func generateAllValidParentheses(depth int) []string {
	soln := make([]string, 0)

	q := queue.New[state](1, state{
		str:   "(",
		left:  depth - 1,
		right: depth,
	})

	for !q.Empty() {
		s := q.DeqHead()

		if s.left > 0 {
			q.EnqHead(state{
				str:   s.str + "(",
				left:  s.left - 1,
				right: s.right,
			})
		}

		// Have to skip left == right cause we have a pattern like ()()
		// If we add a ) to ()(), we get ()()) which is invalid
		if s.right > 0 && s.left != s.right {
			q.EnqHead(state{
				str:   s.str + ")",
				left:  s.left,
				right: s.right - 1,
			})
		}

		if s.left == 0 && s.right == 0 {
			soln = append(soln, s.str)
		}
	}

	return soln
}
