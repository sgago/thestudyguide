package node

type Node[T any] struct {
	Val      T
	Children []*Node[T]
}

func New[T any](val T, children ...*Node[T]) *Node[T] {
	return &Node[T]{
		Val:      val,
		Children: children,
	}
}

func (n *Node[T]) Leaf() bool {
	for _, n := range n.Children {
		if n != nil {
			return false
		}
	}

	return true
}

func (n *Node[T]) AddChild(val T, children ...*Node[T]) {
	n.Children = append(n.Children, New[T](val, children...))
}
