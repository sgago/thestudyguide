package depthfirstsearch

import (
	"fmt"
	"math"
	"testing"

	"sgago/thestudyguide/col/node/binary"
	"sgago/thestudyguide/col/tree/bst"
)

/*
A binary tree is considered balanced if, for every node in the tree,
the difference in the height (or depth) of its left and right
subtrees is at most 1.

In other words, the depth of the two subtrees for every node
in the tree differ by no more than one.

The height (or depth) of a tree is defined as the number
of edges on the longest path from the root node to any leaf node.

Note: An empty tree is considered balanced by definition.

In that case, given a binary tree, determine if it's balanced.

Parameter
tree: A binary tree.
Result
A boolean representing whether the tree given is balanced.

Mmk, so if the difference in height of either left or right subtrees is > 1,
return false.


0|   3
1|	2 4
2| 1   5
*/

func TestBalancedBinaryTree(t *testing.T) {
	tree := bst.New(3, 1, 2, 4, 5, 6)

	bal := balancedBinaryTree(tree)

	fmt.Println(bal)
}

func balancedBinaryTree(tree *bst.Tree[int]) bool {
	left, right := getDepth(tree.Root, 0)

	return int(math.Abs(float64(left-right))) <= 1
}

func getDepth(n *binary.Node[int], lvl int) (int, int) {
	if n == nil {
		return lvl, lvl
	}

	leftMin, leftMax := getDepth(n.Left, lvl+1)
	rightMin, rightMax := getDepth(n.Right, lvl+1)

	return min(leftMin, rightMin), max(leftMax, rightMax)
}
