package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Kth Largest Element in an Array
Find the kth largest element in an unsorted array.
Note that it is the kth largest element in the sorted order,
not necessarily the kth distinct element.

Example 1:
	Input: [3,2,1,5,6,4] and k = 2
	Output: 5

Example 2:
	Input: [3,2,3,1,2,4,5,5,6] and k = 4
	Output: 4
	Note: You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

func TestKLargestElement_K2(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	actual := kLargestElement(nums, 2)
	assert.Equal(t, 5, actual)
}

func TestKLargestElement_K4(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	actual := kLargestElement(nums, 4)
	assert.Equal(t, 4, actual)
}

type MaxHeap []int

var _ (heap.Interface) = (*MaxHeap)(nil)

func NewMaxHeap(nums ...int) *MaxHeap {
	h := &MaxHeap{}
	*h = nums

	heap.Init(h)

	return h
}

// Len implements heap.Interface.
func (m *MaxHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MaxHeap) Less(i int, j int) bool {
	return (*m)[i] > (*m)[j]
}

// Pop implements heap.Interface.
func (m *MaxHeap) Pop() any {
	e := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]

	return e
}

// Push implements heap.Interface.
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m *MaxHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

// kLargest get's the kth largest element in nums using a max heap.
// It heapifies nums in-place to reduce SC to O(1); using a new heap would have a SC of O(N).
// TC to heapify all nums is O(N) and popping all elements is O(KlogN) which means
// our total TC is O(N + KlogN).
//
// ALTERNATE SOLUTION: Sort nums and grab the kth element from the end. SC of O(1) and TC of O(NlogN).
func kLargestElement(nums []int, k int) int {
	// Convert the input slice into a heap immediately using heap.Init
	// This trick reduces space complexity to O(1) since we heapify in-place.
	h := NewMaxHeap(nums...)

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}
