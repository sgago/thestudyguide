package bst

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	tree := New[int]()

	tree.Insert(5, 4, 6, 1, 7)

	result := tree.InOrder()
	fmt.Printf("InOrder: %v\n", result)

	result = tree.PreOrder()
	fmt.Printf("PreOrder: %v\n", result)

	result = tree.PostOrder()
	fmt.Printf("PostOrder: %v\n", result)
}
