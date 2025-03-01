package recoveratreefrompreordertraversal

import (
	"testing"
)

/*
1028. Recover a Tree From Preorder Traversal
https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

We run a preorder depth-first search (DFS) on the root of a binary tree.

At each node in this traversal, we output D dashes (where D is the depth of this node),
then we output the value of this node.  If the depth of a node is D, the depth of its
immediate child is D + 1.  The depth of the root node is 0.

If a node has only one child, that child is guaranteed to be the left child.

Given the output traversal of this traversal, recover the tree and return its root.

Example 1:
	Input: traversal = "1-2--3--4-5--6--7"
	Output: [1,2,5,3,4,6,7]

Example 2:
	Input: traversal = "1-2--3---4-5--6---7"
	Output: [1,2,5,3,null,6,null,4,null,7]

Example 3:
	Input: traversal = "1-401--349---90--88"
	Output: [1,401,null,349,88,90]

Constraints:
- The number of nodes in the original tree is in the range [1, 1000].
- 1 <= Node.val <= 109

======================================================================================

So, we're going to convert the tree from pre-order to level-order traversal.
If a node has one child, then it's guaranteed to be the left child.
We also have to insert nulls if the node has no children. We'll use -1 to represent nulls. This is less obnoxious in golang.
The left most number is the root.
The number of possible, total nodes per level is 2^level. (1, 2 4, 8, 16, 32, ...)
The number of dashes represents the depth.
The numbers can be many digits long.
I wonder if this is just easier to do with a actual binary tree.
*/

func TestRecoverFromPreorder_WithFullTree(t *testing.T) {
	RecoverFromPreorder("0-1--3--4-2--5--6")
}

/*
    0
   1 2
 3  5
7  11
*/

func TestRecoverFromPreorder_WithPartialTree(t *testing.T) {
	RecoverFromPreorder("0-1--3---7-2--5---11")
}

func RecoverFromPreorder(str string) []int {
	i := 0
	n := len(str)
	maxDepth := 0

	for i < n {
		num := 0
		lvl := 0

		for ; i < n && str[i] == '-'; i++ {
			lvl++
		}

		maxDepth = max(lvl, maxDepth)

		for ; i < n && str[i] != '-'; i++ {
			num = num*10 + int(str[i]-'0')
		}

	}

	return []int{}
}
