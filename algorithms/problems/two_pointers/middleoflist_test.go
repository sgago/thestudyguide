package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Middle of a Linked List
https://algo.monster/problems/middle_of_linked_list

Find the middle node of a linked list.
	Input: 0 1 2 3 4
	Output: 2

If the number of nodes is even, then return the second middle node.
	Input: 0 1 2 3 4 5
	Output: 3

======================================================

So, we're trying to return the middle element (if even, return the second element).
In many lang's, we can semi-quickly access the list's length, divide by two, round as needed, grab the element, and call it a day.
Here, however, the problem seems to suggest a list without a known length a head of time, one backed by ptrs to the next element
instead of a slice.

So, what do we do? We could have a slow (s) and fast (f) pointer strategy. Maybe slow and fast move by 1 and 2 nodes, respectively.

s
0 1 2 3 4 5
  f

  s
0 1 2 3 4 5
      f

    s
0 1 2 3 4 5
          f

      s (2nd middle, per problem statement)
0 1 2 3 4 5
          f (end)

We still need to travel n nodes to find length and even/odd.

Yeah, we need to keep the slow pointer so we don't have to iterate all the way back in the list, I see.
This also counts the elements since we need to return slow ptr + 1 to get the right element for an even
amount of elements.
*/

type listNode struct {
	val  int
	next *listNode
}

func TestMiddleOfListWithOddElementCount(t *testing.T) {
	count := 7

	head := &listNode{
		val: 0,
	}

	prev := head

	// Init our ptr-backed list
	for i := 1; i < count; i++ {
		next := &listNode{
			val: i,
		}

		prev.next = next
		prev = next
	}

	elem := findMiddle(head)

	assert.Equal(t, count/2, elem.val)
}

func TestMiddleOfListWithEvenElementCount(t *testing.T) {
	count := 6

	head := &listNode{
		val: 0,
	}

	prev := head

	// Init our ptr-backed list
	for i := 1; i < count; i++ {
		next := &listNode{
			val: i,
		}

		prev.next = next
		prev = next
	}

	elem := findMiddle(head)

	assert.Equal(t, count/2+1, elem.val)
}

func findMiddle(head *listNode) *listNode {
	if head == nil {
		return nil
	}

	count := 1
	slow := head
	fast := head

	for fast.next != nil {
		slow = slow.next // Move slow 1 node right
		fast = fast.next
		count++

		if fast.next != nil {
			fast = fast.next
			count++
		}

		println("slow:", slow.val, "fast:", fast.val)
	}

	if count%2 == 0 {
		// Even amount of elements, so move slow up one node
		slow = slow.next
	}

	return slow
}
