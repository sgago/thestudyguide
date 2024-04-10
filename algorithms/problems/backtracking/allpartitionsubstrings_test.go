/*
Mmk mmk mmk mmk mmk.
Where'd we go wrong here?
I totally created the wrong solution for the wrong problem here.

Step 1: Learn.
	- The problem statement from the start is always a slog.
	- We need absolutely intimacy with all aspects of the problem.
	- Go thru your own inputs.
	- Always rerun the problem's inputs, processing, and outputs with your own inputs.
Step 2: Brainstorm.
    - Create possible solutions.
	- State the N! way, the N^2 way, etc. Output stuff.
	- Watch out for addition/subtraction, trees, graphs, and other tricks.
	- List pitfalls and gotchyas.
Step 3: Psuedocode.
	- Fixing comments is faster than fixing code, always.
	- You'll probably have some mistakes in it. That's cool.
	- If it's DFS, BFS, or DP, time to draw out the tree
Step 4: Code it.

Let's practice now.

282. Expression Add Operators
https://leetcode.com/problems/expression-add-operators/description/

Given a string num that contains only digits and an integer target,
return all possibilities to insert the binary operators '+', '-', and/or '*'
between the digits of num so that the resultant expression evaluates to the target value.

Note that operands in the returned expressions should not contain leading zeros.

Example 1:
	Input: num = "123", target = 6
	Output: ["1*2*3","1+2+3"]
Explanation: Both "1*2*3" and "1+2+3" evaluate to 6.

Example 2:
	Input: num = "232", target = 8
	Output: ["2*3+2","2+3*2"]
Explanation: Both "2*3+2" and "2+3*2" evaluate to 8.

Example 3:
	Input: num = "3456237490", target = 9191
	Output: []
Explanation: There are no expressions that can be created from "3456237490"
to evaluate to 9191.

Constraints:
- 1 <= num.length <= 10
- num consists of only digits.
- -231 <= target <= 231 - 1

LEARN
So, given a number as string, insert addition, subtraction, or multiplication
wherever necessasry to hit some target value.
Say, given 248 to hit 14 -> we can do 2+4+8. This example is lame cause it doesn't have multiple solutions.
Say, given 268 to hit 64 -> we can do (2+6)*8. Again, lame example without mixed operators.
We're not allowed to rearrange the numbers in the string.
We need to watch out for operator precedence. mul > add/sub.
Keywords: return all. combination/permutation.
For each pair of numbers, we have 3 possible operators. total of len(num)-1 * 3 options. For 3 numbers, there's 2*3 => 6 possible solutions.

BRAINSTORM
Brute force also exists. We can loop, keeping a total, find the soln that way. Ehhh, maybe not. This'll get gross fast.

Now, brainstorm some ways to start solving.
We're doing a combination/permutation problem. We have to use all the numbers, we cannot add our own numbers.
We can do a DFS and we don't need to rearrange numbers. So we'll create a node, see if it's a soln, expand it, see if soln, if not trim.
Once we find a solution though, we need to keep going and find all.
We may need to do some memoization, but let's skip until we know it's needed.
Multiplication is realllly tricky. 2*3+2 and 2+3*2 both equal 8, but operator precedence matters significantly.
We are not permitted to evaluate anything until we have hit the final state.

*/

package backtracking

import (
	"testing"

	"sgago/thestudyguide/col/queue"
	stringutil "sgago/thestudyguide/utils/strings"
)

func TestPartition(t *testing.T) {
	partition("mom")
}

type myPartition struct {
	val       string
	remaining string
}

func partition(s string) []string {
	soln := []string{}
	q := queue.New[myPartition](0)

	for i, c := range s {
		q.EnqTail(myPartition{
			val:       string(c),
			remaining: s[i+1:],
		})
	}

	for !q.Empty() {
		p := q.DeqHead()

		if stringutil.Palindrome(p.val) {
			soln = append(soln, p.val)
		}

		if p.remaining != "" {
			next := myPartition{
				val:       p.val + string(p.remaining[0]),
				remaining: p.remaining[1:],
			}

			q.EnqHead(next)
		}
	}

	return soln
}
