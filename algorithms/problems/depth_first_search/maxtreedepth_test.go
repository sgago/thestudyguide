package depthfirstsearch

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/node/binary"
	"sgago/thestudyguide/col/tree/bst"

	"github.com/stretchr/testify/assert"
)

func Test_MaxDepthOfATree_WithLocalState(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, -2, -3, -4, -5, -6, -7}
	tree := bst.New[int](arr...)

	fmt.Printf("DFS traversal preorder is: %v\n", tree.PreOrder())
	depth := FindMaxDepth_WithLocalState(tree)

	assert.Equal(t, 6, depth)
}

func FindMaxDepth_WithLocalState(tree *bst.Tree[int]) int {
	return findMaxDepth_WithLocalState(tree.Root, -1)
}

func findMaxDepth_WithLocalState(node *binary.Node[int], depth int) int {
	if node == nil {
		return depth
	}

	depth++

	left := findMaxDepth_WithLocalState(node.Left, depth)
	right := findMaxDepth_WithLocalState(node.Right, depth)

	if left >= right {
		return left
	}

	return right
}

// This is our "global" state here.
// We could easily slap this variable into a new tree struct or similar.
var maxVal int = -1

func Test_MaxDepthOfATree_WithGlobalState(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, -2, -3, -4, -5, -6, -7}
	tree := bst.New[int](arr...)

	fmt.Printf("DFS traversal preorder is: %v\n", tree.PreOrder())
	depth := FindMaxDepth_WithGlobalState(tree)

	assert.Equal(t, 6, depth)
}

func FindMaxDepth_WithGlobalState(tree *bst.Tree[int]) int {
	maxVal = -1
	findMaxDepth_WithGlobalState(tree.Root, maxVal)
	return maxVal
}

func findMaxDepth_WithGlobalState(node *binary.Node[int], depth int) {
	if node == nil {
		if maxVal < depth {
			maxVal = depth
		}

		return
	}

	depth++

	findMaxDepth_WithGlobalState(node.Left, depth)
	findMaxDepth_WithGlobalState(node.Right, depth)
}
