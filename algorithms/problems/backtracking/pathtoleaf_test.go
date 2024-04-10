package backtracking

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/node/binary"
	"sgago/thestudyguide/col/queue"
	"sgago/thestudyguide/col/tree/bst"
)

func TestPathToLeaf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, -2, -3, -4, -5, -6, -7}
	tree := bst.New[int](arr...)

	fmt.Printf("DFS traversal preorder is: %v\n", tree.PreOrder())
	actual := PathToLeaf(tree)

	fmt.Printf("Paths are:\n%v\n", actual)
}

func PathToLeaf(tree *bst.Tree[int]) *[][]int {
	output := make([][]int, 0, 1)
	path := queue.New[int](3)

	pathToLeaf(tree.Root, path, &output)

	return &output
}

// The queue (q) pointer and output (o) pointers, are similar to passsing global state around.
// IF there's too many args, time to consider either looping + a datastruct or a shared state struct to pass thru to each node.
func pathToLeaf(n *binary.Node[int], q *queue.Queue[int], o *[][]int) {
	if n == nil {
		return
	}

	q.EnqTail(n.Val)

	if n.Left == nil && n.Right == nil {
		// Leaf node is a solution
		// Be careful and actually deep copy the underlying stack slice to output
		// We need to reverse the path slice, it's a stack so the order is flipped, should just use a queue instead...
		soln := q.Slice()
		cpy := make([]int, len(soln))
		copy(cpy, soln)
		*o = append(*o, cpy)
	}

	// Recurse on children
	pathToLeaf(n.Left, q, o)
	pathToLeaf(n.Right, q, o)

	// Pop cause this node is done.
	q.DeqTail()
}
