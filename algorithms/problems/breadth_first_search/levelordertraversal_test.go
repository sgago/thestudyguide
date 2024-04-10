package breadthfirstsearch

import (
	"fmt"
	"testing"

	"sgago/thestudyguide/col/node/binary"
	"sgago/thestudyguide/col/queue"
	"sgago/thestudyguide/col/tree/bst"
)

func TestLevelOrderTraversalBottomUp(t *testing.T) {
	tree := bst.New(5, 4, 6, 3, 7)

	levelOrderTraversalBottomUp(tree)
}

type state struct {
	n   *binary.Node[int]
	lvl int
}

func levelOrderTraversalBottomUp(tree *bst.Tree[int]) [][]int {
	levels := make([][]int, 0)

	q := queue.New[state](7, state{
		n:   tree.Root,
		lvl: 0,
	})

	for !q.Empty() {
		curr := q.DeqHead()

		if len(levels) < curr.lvl+1 {
			levels = append(levels, make([]int, 0))
		}

		levels[curr.lvl] = append(levels[curr.lvl], curr.n.Val)

		fmt.Printf("lvl%d|%d\n", curr.lvl, curr.n.Val)

		if curr.n.Left != nil {
			q.EnqTail(state{
				n:   curr.n.Left,
				lvl: curr.lvl + 1,
			})
		}

		if curr.n.Right != nil {
			q.EnqTail(state{
				n:   curr.n.Right,
				lvl: curr.lvl + 1,
			})
		}
	}

	fmt.Println(levels)

	return levels
}

// This is a bonus dfs soln for fun
func TestLevelOrderTraversalTopDown(t *testing.T) {
	tree := bst.New(5, 4, 6, 3, 7)
	soln := make([][]int, 0)
	levelOrderTraversalTopDown(tree.Root, 0, &soln)

	fmt.Println(soln)
}

func levelOrderTraversalTopDown(n *binary.Node[int], lvl int, soln *[][]int) {
	if n == nil {
		return
	}

	if len(*soln) < lvl+1 {
		*soln = append(*soln, make([]int, 0))
	}

	(*soln)[lvl] = append((*soln)[lvl], n.Val)

	if n.Left != nil {
		levelOrderTraversalTopDown(n.Left, lvl+1, soln)
	}

	if n.Right != nil {
		levelOrderTraversalTopDown(n.Right, lvl+1, soln)
	}
}
