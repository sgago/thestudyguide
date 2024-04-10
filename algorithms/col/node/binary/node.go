package binary

// Node is a binary node with a value and pointers to left and right child nodes.
type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func New[T any](val T) *Node[T] {
	return &Node[T]{
		Val: val,
	}
}

func (b *Node[T]) Leaf() bool {
	return b.Left == nil && b.Right == nil
}
