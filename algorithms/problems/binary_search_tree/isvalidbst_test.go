package binarysearchtree

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/tree/bst"
)

/*
	Valid Binary Search Tree
	https://algo.monster/problems/valid_bst

	A binary search tree is a binary tree with the property that for
	any node, the value of this node is greater than any node in its
	left subtree and less than any node's value in its right subtree.
	In other words, an inorder traversal of a binary search tree yields a
	list of values that is monotonically increasing (strictly increasing).

	Given a binary tree, determine whether it is a binary search tree.

	Notes:
	- Mmk, cool, so what makes a binary search tree?
	- left < current < right, of course.
	- In order traversal should yield monotonically increasing collection.
	- N^2 is make array of in order nodes, then see if it's sorted or whatever.
	- We can simply quit the loop immediately if we find something out of order.
	- For each new node visited key must be > the previous value.
	- Ahh we don't need to keep all values, just the last max value in our state.
*/

func TestValidBinarySearch(t *testing.T) {
	tree := bst.New(8, 6, 7, 5, 3, 0, 9)
	fmt.Printf("%v", tree.InOrder())
}
