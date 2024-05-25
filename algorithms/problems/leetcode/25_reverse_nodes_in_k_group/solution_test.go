package reversenodesinkgroup

/*
25. Reverse Nodes in k-Group
https://leetcode.com/problems/reverse-nodes-in-k-group/

Given the head of a linked list, reverse the nodes of the list k at a time,
and return the modified list.

k is a positive integer and is less than or equal to the length of the
linked list. If the number of nodes is not a multiple of k then left-out nodes,
in the end, should remain as it is.

You may not alter the values in the list's nodes,
only nodes themselves may be changed.

Example 1:
	Input: head = [1,2,3,4,5], k = 2
	Output: [2,1,4,3,5]

Example 2:
	Input: head = [1,2,3,4,5], k = 3
	Output: [3,2,1,4,5]

Constraints:
- The number of nodes in the list is n.
- 1 <= k <= n <= 5000
- 0 <= Node.val <= 1000

Follow-up: Can you solve the problem in O(1) extra memory space?

The Definition for singly-linked list is:
type ListNode struct {
	Val int
	Next *ListNode
}
*/

/*
So, it looks like we're trying to reverse elements in a list
k groups at a time.

For 1 2 3 4 5 and k = 2, it's 2 1 4 3 5. We can't reverse 5
cause the last 5-only group is too small.

For 1 2 3 4 5 and k = 3, it's 3 2 1 4 5. We can't reverse 4 5
cause the last 4 5 group is too small.

For 1 2 3 4 5 and k = 4, it's 4 3 2 1 5. We can't reverse 5,
cause the last 5-only group is too small.

This is a list, so we only have the head node to begin with.
We do not know the length of the list, we only have the head node.

What are our options?
Brute force would be like IDK, using a single pointer to walk
through the list, find a group, then start reversing it.
We could also brute force out pointers into an array and swap those.
The problem want's O(1) mem, so this is suboptimal then.

DFS, DP, none of that feels good at all here or isn't relevant.

Two pointer moving in the same-ish direction? Fast and slow?
Is that it? Draw it out!

1 2 3 4 5 and k = 2.

Start:
f
1 2 3 4 5
s

Group 1 found:
  f
1 2 3 4 5
s

Slow pointer swaps it and advances
  f
2 1 3 4 5
  s

We can advance both pointers or whatever:
    f
2 1 3 4 5
    s

Fast pointer continues to group 2
      f
2 1 3 4 5
    s

Slow pointer swaps it and advances
      f
2 1 4 3 5
      s

Fast pointer can't collect enough elements a 3rd group
        f
2 1 4 3 5
      s

Then quit.

1 2 3
2 1 3
2 3 1

Specifically:
1. Fast pointer finds a group. If not, quit.
2. We need two pointers to do the swapping.

Ahh, this is fun. We have two pointers moving in the same direction.
Then we have two pointers moving in opposite directions to perform the swapping.
The swappers will swap the outside elements, then both move towards each other,
until they are the same or pass each other up. I'll call F the fast pointer
and L and R the swapper pointers.

Yeah, this is close to the proposed solution. TC is
F navigating through the list to drop off the L and R pointers for swapping
is N. Swapping will be N. Like O(N * N) == O(2N) == O(N). Yay.

Mem is O(1).
*/
