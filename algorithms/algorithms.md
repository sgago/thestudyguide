# Algorithms and data structures

- [Algorithms and data structures](#algorithms-and-data-structures)
- [Math](#math)
- [Runtimes](#runtimes)
  - [Summary](#summary)
  - [O(1)](#o1)
  - [O(logN)](#ologn)
  - [O(KlogN)](#oklogn)
  - [O(N)](#on)
  - [O(KN)](#okn)
  - [O(N + M)](#on--m)
  - [O(V + E)](#ov--e)
  - [O(NlogN)](#onlogn)
  - [O(N^2)](#on2)
  - [O(2^N)](#o2n)
  - [O(N!)](#on-1)
  - [Amortized](#amortized)
  - [Time complexity math](#time-complexity-math)
  - [Tricks](#tricks)
- [Hash functions and maps](#hash-functions-and-maps)
  - [Common hash functions](#common-hash-functions)
  - [Pigeonhole principle](#pigeonhole-principle)
  - [References](#references)
- [Sorting](#sorting)
  - [Summary](#summary-1)
  - [Go sort](#go-sort)
  - [Cycle sort](#cycle-sort)
    - [Semi sort](#semi-sort)
- [Binary search](#binary-search)
  - [Go binary search](#go-binary-search)
  - [Binary search keywords](#binary-search-keywords)
- [Trees](#trees)
  - [Terminology](#terminology)
  - [Binary trees](#binary-trees)
  - [Binary search trees](#binary-search-trees)
  - [Balanced and unbalanced binary trees](#balanced-and-unbalanced-binary-trees)
  - [Tree traversal](#tree-traversal)
  - [Tree keywords](#tree-keywords)
- [Heaps](#heaps)
  - [Go heap](#go-heap)
- [Depth first search](#depth-first-search)
  - [Backtracking](#backtracking)
  - [Pruning](#pruning)
- [Graphs](#graphs)
- [Dynamic programming](#dynamic-programming)
  - [Strategy for DP](#strategy-for-dp)
  - [Strategy for looping](#strategy-for-looping)
  - [Using bitmasks for memos](#using-bitmasks-for-memos)
- [Disjoint union set (DSU)](#disjoint-union-set-dsu)
- [Intervals](#intervals)
- [DP or DFS + memoization + pruning?](#dp-or-dfs--memoization--pruning)
- [Trie](#trie)
- [Keywords](#keywords)
- [Concurrency](#concurrency)
  - [Concurrency vs parallelism](#concurrency-vs-parallelism)
  - [Goroutines](#goroutines)
  - [Channels](#channels)
  - [Goroutine patterns](#goroutine-patterns)
    - [Generator pattern](#generator-pattern)
    - [Multiplex pattern](#multiplex-pattern)
  - [Mutexes](#mutexes)
  - [Atomics](#atomics)
    - [Compare-and-swap](#compare-and-swap)

# Math
There's some math to brush up on first. Most of it's high-school level, fortunately!
- **Logarithms:** To what power must we raise a certain base to get a number? We frequently deal with base Log base 2 here, which tells us how many times we need to multiply a number to reach a certain value. For example, `Log2(8) = 3` (`8 = 2 * 2 * 2 = three 2's`) reads like, "How many times do we need to multiple 2 to get 8?". Be aware of context! In mathematics classes, we usually assume Log base 10 when the log base is omitted, but for computers we are *probably* assuming Log base 2!
- **Permutations and factorials:** These come up in combinatorial problems and similar. Say we want to count the number of unique ways we can arrange the letters a, b, and c. There's 6 possible ways (sequences): `3 * 2 * 1` or `3!`. Rarely, does something more complicated come up, but it does happen like with [shuffle sharding](#shuffle-sharding). For example, with a total of 6 items, how many ways can we arrange them in groups of two? This is given by `6!/(2!4!) = 15` or `N!/R!(N - R)!` where `N` is the total elements and `R` is the size of the subsets.
- **Arithmetic sequences** An arithmetic sequence is a sequence of numbers where the difference between consecutive terms is constant. Like `1, 2, 3, 4, 5` or `2, 4, 5, 8, 10`. The sum of an arithmetic sequence can be quickly calculated with `(first_element + last_element) * number_of_elements / 2`.

# Runtimes
How long an algorithm takes to run for a given input. Also called "time complexity" or "TC" for short.

## Summary
Runtime | Name | Example
--- | --- | --- 
`O(1)` | Constant | Math, assignments
`O(alpha(N))` | Inverse Ackermann | Rare. Close to constant time; think O(4) at most. Appears is Disjointed Set Union.
`O(logN)` | Log | Binary search, binary search tree search
`O(KlogN)` | K linear | K binary searches
`O(N)` | Linear | Traverse array, tree traversal
`O(KN)` | K Linear | Performing a linear operation K times
`O(N + M)` | Linear NM | Traverse array, tree traversal on two separate collections
`O(V + E)` | Graph | Traverse graph with V vertices and E edges (dropped vertical bars around V and E cause the make markdown mad)
`O(NlogN)` | Sort | Quick and merge sort, divide n' conquer
`O(N^2)` | Quadratic | Nested loops
`O(2^N)` | Exponential | Combinatorial, backtracking, permutations
`O(N!)` | Factorial | Combinatorial, backtracking, permutations
Amortized | Amortized | High order terms that are rarely done, smaller order done more frequently. Like doing O(N^2) once at startup and O(logN) every other time.

## O(1)
Constant time. A constant set of of operations.
- Hashmap lookups due to pointer arithmetic magic
- Array access
- Math and assignments
- Pushing and popping from a stack

```go
// Some non-looping, non-recursive code
i = 1
i = i + 100000
fmt.Println("more constant time code")
```

## O(logN)
Log time. Grows slowly. log(1,000,000) is only about 20.
- Binary searches, due to constantly splitting solution space in half
- Balanced binary search tree lookups, again cause it's halved.
- Processing number digits
Unless specified, in programmer world, we mean log base 2.

```go
for i := N; i > 0; i /= 2 {
    // Constant time code
}
```

## O(KlogN)
Typically, when you need to do a log(N) process K times. Examples:
- Heap push/pop K times, like merging N sorted lists
- Binary search K times

## O(N)
Linear time. Typically, looping through a data struct a constant number of times. Examples:
- Traversing an entire array or linked list.
- Two pointer solutions
- Tree or graph traversal due to visiting all the nodes or vertices
- Stack and queue

```go
for i := 0; i < N; i++ {
    // Constant time code
}
```

## O(KN)
Typically, when you need to process N K times. Very exciting.

## O(N + M)
Typically, when you have two inputs of size N and M. Say you loop once N times and then loop M times.
Again, very exciting.

## O(V + E)
For both DFS and BFS on a graph, the time complexity is O(|V| + |E|), where V is the number of vertices and E is the number of edges. The number of edges in a graph could be 1 to |V|^2, we really don't know. So we include both terms here.

## O(NlogN)
When we need to do a logN time process N times.
- Divide and conquer, where divide is logN and merge is N
- Sorting can get down to this.

## O(N^2)
Quadratic time. Not terrible for N < 1000, but does grow quickly.
Usually, interviewers want better than this. If you've come up
with a N^2 runtime solution, there's probably something better.
- Nested loops, where outer and inner loops run N times

```go
for i := 0; i < N; i++ {
    for j := 0; j < N; j++ {
        // Constant time code
    }
}

// OR

for i := 0; i < N; i++ {
    for j := i; j < N; j++ { // j is a factor of i, so this is still N^2!
        // Constant time code
    }
}
```

The bottom loop is tricky because `j` is a factor of `i`. So the second example is N^2 TC, too.

## O(2^N)
Grows very rapidly and often requires memoization to reduce runtime.
- Combinatorial problems, backtracking, and subsets
- Often involves recursion

Note, this one is harder to analyze at first.

## O(N!)
Grows insanely rapidly. Only solvable for small N and typically requires memoization.
- Combinatorial problems, backtracking, permutations.

Note, this one is hard to analyze/spot at first.

## Amortized
Amortized time, meaning to gradually write off the initial time costs, if an operation is rarely done. For example, if we had N^2*N tasks, we could consider the solution O(N) instead of O(N^2) = N*O(1) * O(N) if we only do the N^2 task rarely. For example, if we dynamically size an array one time at startup.

## Time complexity math
N is *sort of* like infinity. It swallows smaller N terms and constants. Unlike infinity however, we're trying to indicate something about how long it takes to run a function as N gets larger and larger. N does not swallow up other N values or variables if multiplication is involved, like with NlogN, N^2, KNlogN, etc. N does swallow constants and reduce; for example, 5 * N * N reduces to 5N^2 and N^2 swallows the 5, so the TC is just N^2. Finally, N does swallow smaller factors if addition is used; example is N! + N^2 is just N!. The N! grows much faster than N^2.
- 2N -> N
- N + logN -> N
- NlogN -> NlogN
- 3N^3 + 2N^2 + N -> N^3
- N^2 + 2^N + N! -> N!

Note that, like the other runtimes, we ignore constant factors and lower order terms: 5N^2 + N = 5N^2 -> N^2.

Again, TC is trying to tell us how our function behaves as N increases in size. So, just because N^2 + N reduces to N^2, that N term we dropped still impacts our runtime.

## Tricks
The maximum number of elements, usually noted in the constraints, will give you *clues* about the runtime and *maybe* the corresponding solution. Why? The leetcode runners are docker containers with a very short shelf life by design. There's a maximum number of operations they'll let your run. If your code takes too long, then leetcode will kill your container and return a time limit expired (TLE) error. So, they actually have to limit the number of input elements.

N Constraint | Runtimes | Description
--- | --- | ---
N < 20 | 2^N or N! | Brute force or backtracking.
20 < N < 3000 | N^2 | Dynamic programming.
3000 < N < 10^6 | O(N) or O(NlogN) | 2 pointers, greedy, heap, and sorting.
N > 10^6 | O(logN) or O(1) | Binary search, math, cycle sort adjacent problems.

Again, guessing the optimal solution from the size of input elements constraint is certainly error prone. This could *suggest a maybe solution*.

# Hash functions and maps
In short, a hash function converts arbitrary sized data into a fixed value, typically an Int32. For example, summing all integers in an array and mod'ing them by 100. We convert a bunch of data or text into a smaller, ergonomic number.

Typically, you don't have to write hash functions from scratch outside of, say, for `GetHashcode` for C# objects or similar.

A hash collision occurs when a hash function generates the same hash value for different data.

A good hash function is typically:
- Fast to compute (low time complexity or runtime)
- Very low chance of collision
- All possible values have a somewhat equal chance of occurring.

## Common hash functions
Name | Description
--- | ---
SHA | Cryptographic hash. SHA-3 is the latest. SHA-2 and -1 have vulnerabilities now.
Blake2 | An improvement over SHA-3, high speed and used in crypto mining.
Argon2 | Password hashing function designed to be resistant to brute force or dictionary attacks. Uses large amount of memory (memory-hard) to make attacks more difficult, for hackers using specialized hardware to crack passwords.
MurmurHash | Fast an efficient non cryptographic has. Useful for hash tables.
CRC | Cyclic redundancy check. Non cryptographic. Fast, but not used for security. Typically, the CRC is appended to messages, like HTTP, to check for corruption.
MD5 | Fast 128 bit hash, but no longer recommended for security.

## Pigeonhole principle
Collisions are unavoidable, so we need to design around it. For example, if "anne" and "john" created the same hash, we'd overwrite the same hash table entry. To avoid this, we can use [separate chaining](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining) or other strategies.

## References
[List of hash functions](https://www.geeksforgeeks.org/hash-functions-and-list-types-of-hash-functions/)

# Sorting
- Time complexity = The amount of time it takes to sort the collection as a function of the size of the input data, represented in big O notation. Basic sorting is usually N^2, advanced are usually NlogN.
- Stability = If two elements have equal keys, then the order of these elements remains unchanged. This can be valuable for historical data, user expectations, or multi-criteria sorts where not sorting equal elements is important.
- In-place = The sorting algorithm sorts the input data structure without need to allocate additional memory to store the sorted results. This is valuable for large data sets.
- Simple = Simple algorithms are those that are relatively straightforward to implement with a one or two loops. The more complicated algorithms like quick and merge sort use divide and conquer strategies. This does not mean they are super easy to just bang out, however.
- Adaptable = Sometimes, input data are already somewhat sorted and we can minimize the number of comparisons that we need to make. For example, gnome sort is O(N) for an already sorted collection!
- Parallelizable = The sorting algorithm can divide the the sorting into subtasks that can be executed in parallel. Merge, quick, radix, and bucket sorts are all parallelizable due to divide and conquer.

## Summary
A summary of common algorithms, courtesy of ChatGPT.
| Algorithm | Time Complexity (Worst Case) | Stable | In-Place | Adaptable | Parallelizable | Description |
| --- | --- | --- | --- | --- | --- | --- |
| [Bubble](./algo/sort/bubble/bubble.go) | O(n^2) | Yes | Yes | Yes | Yes (limited) | Repeatedly compares and swaps adjacent elements.          |
| [Selection](./algo/sort/selection/selection.go) | O(n^2) | No | Yes | No | Yes (limited) | Repeatedly selects the minimum element and swaps with the current position. |
| [Insertion](./algo/sort/insertion/insertion.go) | O(n^2) | Yes | Yes | Yes | Yes (limited) | Builds the sorted list one element at a time by inserting into the correct position. |
| [Merge](./algo/sort/merge/merge.go) | O(n log n) | Yes | No | Yes | Yes | Divides the input into halves, recursively sorts them, and merges them. |
| [Quick](./algo/sort/quick/quick.go) | O(n^2) (rare), O(n log n) | No | Yes | Yes | Yes (limited) | Chooses a pivot, partitions the data, and recursively sorts the partitions. |
| Heap | O(n log n) | No | Yes | No | Yes | Builds a binary heap and repeatedly extracts the maximum element. |
| Shell | O(n log^2 n) (worst known) | No | Yes | No | No | A variation of insertion sort with multiple passes and varying gap sizes. |
| Radix | O(nk) (k is the number of digits) | Yes | Yes | No | Yes | Processes digits or elements in multiple passes, each pass sorted independently. |
| Bucket | O(n^2) (worst case) | Yes | No | Yes | Yes | Distributes elements into buckets and sorts each bucket independently. |
| [Cycle](./algo/sort/cycle/cycle.go) | O(N^2) | No | Yes | No | No | Based on the idea that elements can be sorted by setting them to their correct index position. Theoretically optimal in terms of writes.

## Go sort
We don't typically author sort algorithms from scratch in production. Some computer scientist has implemented pattern defeating quick sort (pdqsort) for us so that we don't have to. In go, use `slices.Sort` or similar:
```go
arr := []int{5, 3, 1, 4, 2}
slices.Sort(arr)

// OR

cmp := func(a, b int) int {
  if a == b {
    return 0
  } else if a > b {
    return 1
  }

  return -1
}

slices.SortFunc(arr, cmp)
```

## Cycle sort
We know a collection is sorted if all elements are in increasing or decreasing order, but how would we know if an individual element is sorted? Naively, we might think `left element < the element being considered < right element` but this won't work for a collection like `5, 1, 2, 3, 4`. We know an element is sorted if it's index is equal to a count of all elements smaller than it. For example, in `1, 2, 3, 4, 5`, we know 2 is sorted because its index in the slice is equal a count of all numbers less than it, in this case, just 1.

[Cycle sort](./sort/cycle/cycle.go) is based on this idea of elements being at their proper index position.

### Semi sort
There's a class of problems that are solved via semi-sorting elements. (Warning: This is a concept I'm making up - it's not formal or anything and might have a better name somewhere.) For example, [find the minimum missing positive number](https://leetcode.com/problems/first-missing-positive/description/) like problems can be solved this way. (Another warning: Semi sort uses O(N) space and it isn't an optimal solution for this problem!) Finding the missing number can be solved faster than O(NlogN) time if we only partially, kind-of sort the array. To do this, we simply move elements to the same index as their value ([here](./sort/semi/semi.go)) in a new slice. That is, we copy 0 to index position 0 in the new slice, 1 to index position 1, etc. "But what about index out of range errors!?" If element values is not in-range, we can just ignore them or use a bad value.

# Binary search

```go
func binarySearch(nums []int) {
  ans := 0
  left, right := 0, len(nums)-1

  for left <= right {
    mid := left + (right-left)/2

    if nums[mid] > nums[left] {
      // Pull in left side
      left = mid + 1
    } else if nums[mid] <= nums[right] {
      // Pull in right side
      right = mid - 1
      ans = mid // Maybe get answer from here
    }
  }

  return ans
}
```

## Go binary search
For a bounded binary search, one where we have a sorted collection and an explicit target, we can use `slices.BinarySearch` of `slices.BinarySearchFunc` like

```go
func TestBoundedBinarySearch(t *testing.T) {
	needle := 3
	haystack := []int{1, 3, 5, 7, 9}

	// Searching for a value in a sorted slice
  // Output is 1, true
	idx, ok := slices.BinarySearch(haystack, needle)
}
```

## Binary search keywords
These keywords may indicate a binary search if they appear in the story problem text.
- Sorted integers, sorted strings
- Rotated array, find peak
- Any array/slice that could be considered implicitly sorted

# Trees
Trees are a type of graph composed of nodes and edges.
- Trees are acyclic, nodes don't loop back to themselves and create cycles.
- There's a path from the root node to any other node.
- Trees have N-1 edges, when N is the number of nodes.
- Nodes have exactly one parent node.
- Trees are directed.
- Trees are rooted.

## Terminology
Name | Description
--- | ---
Root node | The top most ancestor node, one with out any parents.
Internal node | Every node that has at least one child.
Leaf node | Every node that does not have any children.
Ancestor | All nodes between the path from the child to the parent.
Descendent | All nodes between the path from a parent to the child.
Level | Number of ancestors from the current node to the root.
Arity | Number of operators or terms. For trees, where each node has no more than N-ary children.

I dislike ancestor/descendent. Call 'em parent and child like we're all 5 years old.

Example of a tree:
```
    A
   / \
  B   C
 / \   \
D   E   G
```
- A through G are all nodes.
- / and \ are edges. There's N-1 = 6-1 = 5 edges and 6 nodes.
- A is the root node.
- D, E, and G are all leaf nodes.
- A, B, and C are internal nodes.
- A is at level 0, B and C at level 1, and D, E, and G at level 2.
- A is a parent of B and C. B is a parent of D and E. C is a parent of G. D, E, and G are leaf nodes and have no children themselves.
- D and E are a child of B. G is a child of C. B and C are children of A. A is the root node and has no parents.

## Binary trees
Binary trees are a tree where each node has 0 to 2 children.

A full binary tree is one in which every node has 0 or 2 children. 1 is not allowed.
```
      x
     / \
    x   x
   / \
  x   x
 / \
x   x
```

A complete binary tree is where all the levels, except the last, are completely filled out. In the last level, the nodes are as far left as possible. This shows up in heaps.
```
      x
     / \
    x   x
   / \
  x   x
```

A perfect binary tree is one in which all the internal nodes have exactly 2 children. 1 or 0 are children are not allowed. All leaf nodes have the same number of children.
```
       x
     /   \
    x     x
   / \   / \
  x   x x   x
```
Perfect trees are used to estimate time complexity for combinatorial problems where the search space is a perfect binary tree. They have some unique properties.
- The number of nodes is 2^L-1 where L is the number of levels.
- The number of internal nodes is # of leaf nodes - 1.
- The total number of nodes is = 2 * leaf nodes - 1.

## Binary search trees
Binary search trees (BSTs) are a special type of binary tree where all left descendants < node < all right descendants.
```
      8
     / \
    3   10
   / \    \
  1   5    14
       \
        7
```
Notice that 3 is to the left of 8 because 3 < 8. Similarly, 14 is to the right of 10.
Note that the in-order traversal of the tree visits the nodes in monotonically increasing order.

## Balanced and unbalanced binary trees
Unbalanced trees are have a search time of N. The start to look more like a list than a tree.
```
1
 \
  2
   \
    3
     \
      4
```
Balanced binary trees are those where the difference in height (levels) between the left and right subtrees of all nodes is not more than 1. Balanced trees allow for a search time of logN.
```
      8
     / \
    3   10
   / \    \
  1   5    14
       \
        7
```

Balanced binary trees include [red-black](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) and [AVL](https://en.wikipedia.org/wiki/AVL_tree) trees.

## Tree traversal
Tree traversal are types of traveling through the nodes of a tree.
- In-order visits the left branch, current node, then right branch.
- Pre-order visits the current node, left subtree, and right subtree.
- Post-order visits the left subtree, right subtree, and current node.

For example, given the following binary search tree,
```
     8
    / \
   3   10
  / \    \
 1   5    14
      \
       7
```
the visits to each node would be:
- In-order: 1 3 5 7 8 10 14
- Pre-order: 8 3 1 5 7 10 14
- Post-order: 1 7 5 3 14 10 8

## Tree keywords
These keywords may indicate a tree if they appear in the story problem text.
- Shortest, level-order, zig-zag order, etc.
- Max depth

# Heaps
A min heap is a special tree data structure where
1. Almost complete - every level in the tree is almost filled, except the last level. The last level is left justified.
1. Each node has a greater key (priority) than it's parent.

Here's a heap structure:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
  |
  12
```

This is not a heap:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
  |
  5 // Need to heapify this, 5 < 7, so our heap property is broken
```

Also, not a heap:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
      |
      12 // Need to left justify this
```

Also, not a heap:
```
       1
     /   \
    2     3
   / \   /
  7   8 9 // Need to fill intermediate levels
  |
  12
```

Couple notes:
- Priority queue is an abstraction over a heap, minheap and maxheap are the concrete implementations.
- Max heaps are the same, but we just change the each-child-key-is-greater property to each child is less.
- Usually, heaps are binary tree, but you can also get k-ary heaps or k-heaps.
- A priority queue is an abstraction on the heap, a min/max heap is the concrete implementation.
- They are typically implemented with an array.
- It's a sorted tree but it is not a binary search tree.

Heaps support three main operations:
1. Heapify - Rearrange the nodes such that the heap such that the nodes are in min keys are always the at the root. This is an O(logN) operation.
2. Insert - Inserts a new element into the heap and calls heapify because to maintain the heap properties. This is an O(logN) operation.
3. Pop/Delete - Removes and returns the min root element. Another O(logN) operation.

In it's use, it's sort of like a stack. Push nodes in, pop nodes out except you always get the min keyed node.

## Go heap
Like sorting, we don't have to author heaps from scratch in go.
For example, we can alias a slice type and implement heap.Interface on it:
```go
// Alias a slice type
type MinHeap []int

// And implement the heap.Interface type.
var _ heap.Interface = (*MinHeap)(nil)

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Pop() any {
	result := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return result
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func TestHeap(t *testing.T) {
	// Create an empty MinHeap
	h := &MinHeap{}

	// Push elements onto the heap h, using heap.Push(h, whatever)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 2)

	// Pop elements from the heap h (retrieve in sorted order for a min-heap)
	// using the heap.Pop(h)
	for h.Len() > 0 {
		fmt.Printf("%d\v", heap.Pop(h))
	}
}
```

# Depth first search
A depth first search (DFS) looks for solutions by going deep first. That is, it searches for solutions in a pre-order traversed way. Some more terminology:
- Backtracking - returning after visiting a non-solution node
- Divide and conqueror - When we have two or more recursive calls, that splits our issue into subproblems, like O(logN).

When do we use a DFS? We can use them in
1. Trees
  - Traverse through a tree to find, create, delete nodes
  - Traverse through a tree to find the max subtree, detect a balanced tree, etc.
2. Combinatorial problems
  - Find all the different ways of arranging something
  - Find all possible combinations of something 
  - Find all solutions to a puzzle
3. Graphs
  - Find a path from A to B in a graph
  - Find connected components
  - Detect cycles

There are two main ways of coding a DFS. We can use recursion or we can use a stack + loop.
1. With recursion, we call the recursive func on the next state.
2. With a loop, we pop off a stack, add any possible solutions, and then keep going.

One thing to keep in mind with DFS, is kicking state into and out of the recursive functions and/or stack.

The template for a DFS via recursion is
```
function dfs(node, state):
    if node is null:
        ...
        return

    left = dfs(node.left, state)
    right = dfs(node.right, state)

        ...

    return ...
```

The main hard parts of DFS are 1. deciding which state to pass in, 2. deciding which state to pass out and 3. sorting out the recursive calls.

Now, sometimes, when solving DFS problems, we need to
- Pass information back up through the return value, like the max depth.
- Pass information about state into the recursive calls, like max value
Alternatively, we can store state, say a max number, in a global variable.

## Backtracking
Backtracking tacks on some new concepts on top of trees and DFS. In short, the problems don't always give us tree or graph to work with. Sometimes, we have to make the generate the tree or graph as we go.
1. Again, we can make the tree as we go, creating and deleting child nodes as we traverse through.
2. We can drag state around via parameters/returns or with global/struct variables.
3. If we get some crazy 2^N to N! memory usage with backtracking combinatorial problems, we may need to memoize intermediate solutions to cut down on memory usage. For a small to mid N, N! will kill our poor computer. Memoize typically means using a map or similar to store intermediate and final solutions to the combinatorial problems.

## Pruning

# Graphs
Trees are rooted, connected, acyclic, undirected graphs. Trees contain N nodes and N-1 edges and there are only one path between 2 nodes.

This is a tree (and a graph):
```
    1
   / \
  2   3
 /
4
```

This is only a graph due to the cycle and disconnected vertex.
```
    1
   / \
  2 - 3   // Cycle among nodes 1, 2, and 3

4         // Node 4 is disconnected from others
```

Trees and graphs have different terminology
- Vertices are nodes in trees.
- Vertices are connected by edges.
- Vertices connected by an edge are neighbors (children and parents in trees).
- Edges can be directed or undirected. Usually the edges are undirected.
- Paths are sequences of vertices. Cycles start and end at the same vertex.
- A connected graph means every vertex is joined by a path to a vertex; otherwise, the graph is disconnected.

Typically, graphs are stored via adjacency lists or maps. For example, this graph
```
    1
   / \
  2   3
 /
4
```

can be represented in go via a map like

```go
graph := map[int][]int{
  1: {2, 3},
  2: {1, 2, 3},
  3: {1, 2},
  4: {2},
}
```

Note that we don't need to have an adjacency list upfront to solve problems.

So, for BFS and DFS on the graph, we can do the stack/queue dance through it. However, we need a way to dodge any cycles or we'll get stuck in an inf loop or stack overflow. To do this, we can store a map of visited nodes when searching. Like

```go
visited := make(map[int]bool)
visited[1] = true
```

Another clever trick is to wipe out the value in the adjacency graph somehow. Make it negative or something that indicates we visited it to avoid storing a whole other data structure for visited/not-visited stuff.

When deciding between BFS or DFS to explore graphs, choose BFS for shortest distance or graphs of unknown or infinite sizes due to exploring all adjacent neighbors first.

DFS is better at using less memory for wide graphs (graphs with large breadth of factors). Put another way, BFS stores the breadth of the graph as it searches. DFS is also better at finding nodes that are far away such as a maze exit.

# Dynamic programming
A problem can be solved via dynamic programming (DP) if
1. The problem can be divided into sub-problems
2. The sub-problems from bullet 1 overlap

To put DP very simply, all you're trying to do is fill up a 1D or 2D array up with solutions to subproblems such that some slice element has the final solution to the entire problem in it. All the other elements will be solutions to subproblems.

Ultimately, when solving DP problems, we're trying to develop the *recurrence relation* and it is simply critical. For example, the recurrence relation to tabulate any Fibonacci number into a `dp` memo is `dp[i] = dp[i - 1] + dp[i - 2]`.

To be blunt, some recurrence relations just aren't obvious. Take finding the longest increasing subsequence (LIS). The recurrence relation from is `lis(i) = max(lis(i-1), lis(i-2) ... lis(0))` but only for `nums at i-1, i-2, ..., 0 < num at i`. One of the easier ways to figure this out is just to write out your `dp` memo for every loop iteration for some example they give you.

For example, when I started to solve the longest increasing subsequence [here](./algorithms/problems/dynamic_programming/bottom_up/longest_increasing_subsequence/longest_increasing_subsequence_test.go), I just started to write out the entire memo and gave myself some notes like:
```
1 1 1 _ 1 1 1 _ 1 1 1  <- Our initial dp memo
1 2 1 _ 1 1 1 _ 1 1 1  <- 0 to 1 is a LIS of 2.
1 2 3 _ 1 1 1 _ 1 1 1  <- 0 to 1 to 3 is a LIS of 3.
1 2 3 _ 3 1 1 _ 1 1 1  <- 0 to 1 to 2 is a LIS of 3.
1 2 3 _ 3 4 1 _ 1 1 1
1 2 3 _ 3 4 5 _ 1 1 1 <- We're done here... we just don't "know" that yet
1 2 3 _ 3 4 5 _ 1 1 1
1 2 3 _ 3 4 5 _ 1 2 1
1 2 3 _ 3 4 5 _ 1 2 4 <- Done. 5 is biggest LIS count with values 0 1 2 4 5 or 0 1 3 4 5

The recurrence relation is like dp[i] = max(dp[i], dp[prev_i] + 1) but only dp[prev_i] where
nums[i] > nums[prev_i], because otherwise it wouldn't be only increasing numbers.
```

The notes are ugly, but they get the job done.

Really, DP is typically similar in efficiency to DFS + memoization + pruning. Pruning is important, to save space and reduced wasted calculations. We typically call DP *bottom-up* and DFS + memoization + pruning *top-down*.

## Strategy for DP
Broadly speaking, if this is your first sojourn into DP-land, the big question you're trying to answer is "How the !@#$ do I fill up this slice such that one of the values has the final answer?"

To begin solving DP problems, we can answer some basic questions:
1. What goes in our memo? It's typically the same thing whatever the problem is asking for. For example, if the problem is asking for counts of something, then we're memoizing counts. If it's asking for maximum values, then we memoize maximum values. You might also need to determine here if you need a 1D or 2D memo.
2. What are the initial conditions? Call this dp[0].
3. What is the absolutely smallest subproblem I can start with?
4. How do I get from dp[0] to dp[1] using that smallest subproblem?
5. If dp[0] to dp[1] doesn't have overlap, then solve dp[1] to dp[2]. Repeat this until you find something with overlap. It's important to know what to do here.

An easier way might be to type out the entire DP memo for 5 to 10 iterations - enough so that you have to work through some overlap. Literally, just type out loops or different problems until it "clicks". If it doesn't click, just keep going.
1. Declare your base or initial conditions. You need to start somewhere. Sometimes it's zero, the first element, etc.
2. Clearly define how you get for your base state to the next state.
3. Again, you need to solve a sub problem that overlaps. If going from base to the next state doesn't have overlap, then you need to go keep solving subproblems until you get some overlap. This will inform how you deal with the overlap.

A more formal set of instructions for solving DP might be:
1. **Identify overlapping subproblems:** Determine if the problem exhibits overlapping solutions or overlapping solutions. If not, DP will not work and you'll need to try something else like DFS or graphs.
2. **Handle or preprocess the input (optional):** In some cases, the problem might not provide a straightforward array for DP. For example, see the [perfect squares problems](./algo/problems/dynamic_programming/top_down/perfectsquares_test.go) where we need to generate all the perfect squares first or the [largest divisible subset problem](.algo/problems/dynamic_programming/top_down/largestdivsubset_test.go) where the input needs to be sorted first. Futhermore, you might not realize this straight away which makes the understanding the problem and psuedocoding the problem important.
3. **Define the memo:** Clarify what needs to memoized. In many cases, the memo is the same as the problem's output. For sequences, we can usually get by with a 1D map or slice. For grid, interval, and dual-sequence problems, we need a 2D map or slice.
4. **Initialize the memo with base cases:** We can typically need to initialize the memo with the base cases. For example, in the [perfect squares problem](./problems/dynamic/perfectsquares_test.go), we can initialize the memo with perfect squares like 1, 4, 9, 25, etc. immediately. Other times, we may need to loop through the 2D map or slice and declare base cases like for the coin game.
5. **Two loops** Typically, you'll get an embedded looping solution (one loop in another).
6. **Define the recurrence relation:** Develop the recurrence relation, a formula for transition from one state to the next. For example, clearly state how to transition from `dp[i]` to `dp[i+1]`. Start with transition from the base cases to the next case. Write the recurrence relation in something *you* can understand for coding like:
```
The recurrence relation is:
memo[i] = max(memo[i-1], memo[i-2]... memo[0])+1
  - applies ONLY for each memo[i-1] when nums[i]%nums[i-1]==0.
  - memo[0] is a base case and equals 0
```
1. **(Optional) Optimize the memo:** Sometimes you can cut down on the amount of memory that the memo consumes. See Bitmasking.

## Strategy for looping

1. **Declare a state struct** that will carry values we need to check for possible solutions and expand the next states, if any.
```go
type state struct {
  count int
  total int
}
```

2. **Initialize a queue with base states**
```go
	for i := 0; i*i <= target; i++ {
		sqr := i * i

		if sqr == target {
			return 1
		}

		memo[sqr] = 1
		sqrs = append(sqrs, sqr)

		q.EnqHead(state{
			count: 1,
			total: sqr,
			str:   fmt.Sprint(sqr),
		})
	}
```

3. **Create a for-queue-not-empty loop**
```go
	for !q.Empty() {
    // TODO: Get next value
    // TODO: Check if it's a solution
    // TODO: Enqueue next set of states, if any
    // TODO: Optionally, add memoizing
	}
```

4. **Dequeue the current state**
```go
	for !q.Empty() {
    // Get next value
    curr := q.DeqHead()

    // TODO: Check if it's a solution
    // TODO: Enqueue next set of states, if any
    // TODO: Optionally, add memoizing
	}
```

5. **Check if current state is a solution**
```go
	for !q.Empty() {
    // Get next value
    curr := q.DeqHead()

    for _, sqr := range sqrs {
      // Check if it's a solution
      next := curr.total + sqr
      count := curr.count + 1

      // Is next a solution?
      if next == target {
        return count
      }

      // Is next > target? If so, there's no solution
			// so continue on to the next value in our queue.
			if next > target {
				continue
			}

      // TODO: Enqueue next set of states, if any
      // TODO: Optionally, add memoizing
    }
	}
```

6. **Enqueue the next state(s), if any, from the current state**
```go
	for !q.Empty() {
    // Get next value
    curr := q.DeqHead()

    for _, sqr := range sqrs {
      // Check if it's a solution
      next := curr.total + sqr
      count := curr.count + 1

      // Is next a solution?
      if next == target {
        return count
      }

      // Is next > target? If so, there's no solution
			// so continue on to the next value in our queue.
			if next > target {
				continue
			}

      // Enqueue next set of states, if any
			q.EnqTail(state{
				count: curr.count + 1,
				total: curr.total + sqr,
			})

      // TODO: Optionally, add memoizing
    }
	}
```

7. **Optionally, add memoization to reduce loop iterations**
```go
	for !q.Empty() {
    // Get next value
    curr := q.DeqHead()

    for _, sqr := range sqrs {
      // Check if it's a solution
      next := curr.total + sqr
      count := curr.count + 1

      // Is next a solution?
      if next == target {
        return count
      }

      // Is next > target? If so, there's no solution
			// so continue on to the next value in our queue.
			if next > target {
				continue
			}

      // Enqueue next set of states, if any
			q.EnqTail(state{
				count: curr.count + 1,
				total: curr.total + sqr,
			})

      // Optionally, add memoizing

      // Have we looked at this next value before?
			// Did we get to the next value with a lower count?
			// If so, there's no point in enq'ing the next round of
			// perfect sqrs. We got to next in some more efficient
			// way, so we can just continue on to the next number.
			if memoCount, ok := memo[next]; ok && count >= memoCount {
				continue
			}

			memo[next] = curr.count
    }
	}
```

See the complete looping solution [here](./algorithms/problems/dynamic_programming/looping/perfect_squares/perfectsquares_test.go).

## Using bitmasks for memos
By using a bitmask, we can replace bool arrays like `make([]bool, 64)` with a single `uint64` value. Pretty sweet, right? To do so, we can use all sorts of clever tricks with bits and bitwise operators `&`, `|`, `^`, `<<`, and `>>`.

Some common bit operations you'll see are:
- `1 << i` left shifts the 1, effectively multiplying the number by 2. Conversely, `>>` rights shift the number which divides it in 2.
- `bitmask & (1 << i)` checks if the ith bit in the bitmask is set or not.
- `bitmask ^ (1 << i)` toggles flips the ith bit, setting it to 1 if it was 0 or 0 if it was 1.

I've authored a bitflag collection to help handle this memory optimization technique [here](./algorithms/col/bitflags/bitflags.go).

# Disjoint union set (DSU)

# Intervals
```
start1-----end1   // Interval 1

  start2-----------end2    // Interval 2 that overlaps with 1

                           start3-----end3    // Interval 3 that doesn't overlap with either 1 or 2
```
We can determine if intervals 1 and 2 overlap if `end1 >= start2 && end2 >= start1`. Notice that the formula returns false for intervals 1 and 3.

# DP or DFS + memoization + pruning?
Classically, DP problems can be solved in either top-down or bottom-up approach. The author, making stuff up again from practicing, defines three ways of solving these:
1. **Recursion** - A DFS using recursion with memoization, pruning, and backtracking. The classic definition of top-down.
2. **Looping** - A DFS using a stack or queue to track state instead of passing values around via function calls as we do with recursion. Uses memoization to skip inserting less efficient solutions into the stack or queue. Another top-down solution, similar to above but different.
3. **1D or 2D slice/map** - We solve subproblems and store answers into a 1D or 2D slice/map (aka dynamic programming or DP). If there's overlap in the index, we take the "better" solution. The final location in the slice or map will hold the answer. Again, this is the classic method for bottom-up approaches.

Honestly, if this was a production need, I would say it depends on the problem and how much we care about performance, readability, etc.

But this is "interview programming". Therefore, you might consider simply picking top-down or bottom-up based on comfort level alone.

Generally, from my personal practice, I prefer looping. Why?
- After solving n-million DFS problems with recursion I always end up with a public and private function. The private function that actually does the DFS has an issue with tons of telescoping parameters and it's unclear which of these are inputs and outputs. Invariably, I want to change the order, add a complex type, or simply get confused about which parameters are for what. Eventually, you just decide to slap all the parameters into a state struct because it's easier to modify and use. No need to worry about the order struct fields appear in!
- After solving n-million DP problems, the recurrence relation, the formula that gets you to dp[next] from dp[previous states], is hard to spot yet critical to figure out. Let me put this another way. **You 100% need to develop the recurrence relation formula for getting from a previous state to the next state. You need to do this in a reasonable amount of time. If you have the recurrence relation, you have almost everything you need. If you don't have the recurrence relation, you have nothing. If you don't have the recurrence relation, don't bother starting to code anything. This is all or nothing development - you either have the whole thing or you don't.** Again, you will 100% absolutely need to type out the the entire dp memo as it loops via comments, starting with the initial conditions. And solving one or two subproblems is just not enough to see the solution in its entirety. You need to type out like 5 to 10 iterations and learn how to solve problems that overlap. If you don't type it out, the recurrence relation will be hard to see.

Again, top-down looping feels like it lends itself to iterative problem solving. It feels consistent even though it might not always be faster than DP. With DP, you're going to get a unique recurrence relation and memo every time.
1. Read the problem, determine WTF is going on.
2. Create a `state` struct to track things we need to expand/solve the problem.
3. Add a queue and load up it's initial state.
4. Write a for-queue-not-empty loop.
5. Pop an item off the queue and check for solutions, if any.
6. Create and push the next states into the queue, if any. We're going to make every combination of everything every time.
7. Once that's all solved, add memoization as an optimization. *Typically*, you can memoize the solution but this not always true! For example, see the [minimum XOR sum of two arrays solution](./algorithms/problems/leetcode/1879_minimum_xor_sum_of_two_arrays/) where there's no value we can store. *Usually*, you can memoize. If you can use a memo, clarify what needs to memoized. In *most* cases, the memo is the same as, or aligns closely with, the problems output.

But hey, again, that's just my experience solving recursion, looping, and DP problems. Your mileage may vary.


# Trie
Tries, pronounced like "tries" to distinguish it from "trees", is a k-arity tree used as a lookup structure for prefixes for autocomplete and prefix count type problems. Tries do not typically store their key. Typically they are used with characters and strings, but they can also be used with bits or numbers.

![](./diagrams/out/trie-example/trie-example.svg)

A simple Trie implementation is [here](./algo/col/trie/trie.go).

# Keywords

# Concurrency
Here we give a deep treatment of concurrency in go: goroutines, locks, atomics, compare-and-swap, etc.

[These slides](https://go.dev/talks/2012/concurrency.slide#1) by Rob Pike have a good overview and introduction into gorountines, channels, and patterns.

## Concurrency vs parallelism
Concurrency and parallelism are two confused terms in the software space. For this section, they are not the same.
- *Concurrency* is the composition of independently operating computations. A single core machine can be concurrent but not parallel. A single core machine can rapidly context switch between a word processor and music player, giving an illusion of parallelism but it's actually not. Concurrency is a broader concept and does not necessarily imply actual simultaneous execution.
- *Parallelism* is actual simultaneous execution of multiple executions through multiple cores. For instance, in a multi-core machine, we could run the music player and word processor on individual cores. (Or you could run them both on a single core through concurrency...)

Clear as mud? Good.

## Goroutines
Let's start with a boring func first:
```go
func boring(msg string) {
  for i := 0; ; i++ {
    fmt.Println(msg, i)
    time.Sleep(time.Second)
  }
}
```

In effectively an infinite loop, it prints a msg and sleeps. To make things more exciting for learning concurrency, let's give it a random amount of time.
```go
func boring(msg string) {
  for i := 0; ; i++ {
    fmt.Println(msg, i)
    time.Sleep(time.Duration(rand.Intn(1e3))*time.Second)
  }
}
```

We can invoke this boring func via:
```go
func main() {
  boring("kind of boring")
}
```

Now, we can launch a goroutine by using the go keyword.
```go
func main() {
  go boring("kind of boring")
}
```

Unfortunately, `go boring` is stopped almost immediately when main returns and the program stops. We can let the goroutine hang around a little bit by sleeping the main thread.

```go
func main() {
    go boring("boring!")
    fmt.Println("I'm listening.")
    time.Sleep(2 * time.Second)
    fmt.Println("You're boring; I'm leaving.")
}
```

In short, you can think of goroutines as cheap threads. Under the hood, they most definitely are not threads though they use threads. Similar to a thread, each goroutine has its own stack and are dynamically multiplexed onto threads as needed to keep them all running.

## Channels
In the prior examples, main and boring never actually communicated with one another. We can let two goroutines communicate via a channel.

```go
var ch chan int
ch = make(chan int)
// OR
ch := make(chan int)

ch <- 1 // Send a value
val := <-ch // Get a value
```

## Goroutine patterns

### Generator pattern
A generator pattern creates a hidden goroutine and a channel to communicate with.

```go
// Channels are first-class citizens so we can pass them around like integers. This generate function returns an output-only channel that we can use for communication and synchronization.
generate := func(msg string) <-chan string {
  // Make a chan string and give it to the caller so it can talk to us
  ch := make(chan string)

  // Fire off an internal goroutine
  go func() {
    for i := 0; ; i++ {
      ch <- fmt.Sprint(msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()

  return ch // And give the channel back to the caller, so we can talk
}

// And, elsewhere, use the generate function like this:
out := generate("We're talking!")

for i := 0; i < 5; i++ {
  // Get a value out of the chan 5 times, neat.
  // Warning: This blocks until generate outs a value!
  // Both sides of a channel need to be ready.
  fmt.Println(<-out)
}

/* Output:
We're talking!0
We're talking!1
We're talking!2
We're talking!3
We're talking!4
*/
```

### Multiplex pattern
The multiplex or fan-in pattern merges multiple channel values into one channel. This pattern builds off of the generate pattern.

```go
// This multiplex functions will join together the input channels.
multiplex := func(ch1, ch2 <-chan string) <-chan string {
  c := make(chan string)

  go func() {
    for {
      c <- <-ch1
    }
  }()

  go func() {
    for {
      c <- <-ch2
    }
  }()

  return c
}

// And, elsewhere, use the generate + multiplex functions like this:
joe := generate("joe")
ann := generate("ann")

out := multiplex(joe, ann) // The one output channel to rule them all.

for i := 0; i < 10; i++ {
  // The fan in pattern has unblocked joe from ann and vice versa.
  // This speeds up the output but we lose ordering.
  // Maybe ordering doesn't matter for your solution or maybe it does.
  fmt.Println(<-out)
}

/* Output:
joe0
ann0
joe1
joe2
ann1
joe3
ann2
ann3
joe4
ann4
*/
```

## Mutexes
Go's `sync` package supports mutexes.

See an embedded struct RW locker [here](./algorithms/col/cas/stack/stack.go) and using it to create a concurrent stack [here](./algorithms/col/stack/stack.go).


## Atomics
Go also supports low-level atomic primitives in the `sync/atomic` package. This can be used for thread-safe primitives like

```go
	// Declare an atomic counter like:
	var cnt atomic.Int32

	// Use the atomic counter cnt like
	cnt.Store(3)
	fmt.Println("cnt:", cnt.Add(3)) // cnt: 6
	fmt.Println("cnt:", cnt.Load()) // cnt: 6
```

### Compare-and-swap
The `sync/atomic` package also supports compare-and-swap. We can use this to create lock-free data structures.

```go
	// Compare-and-swap (CAS) can feel a little goofy
	// if it's the first time you're seeing it.

	// CAS steps:
	// 1. Load cnt via cnt.Load()
	// 2. Compare current cnt value to the old value from cnt.Load()
	// 3. If they're the same, cnt.CompareAndSwap swaps the new value in.
	//    However, if cnt's current value does not match the old value,
	//    CompareAndSwap returns false. In that case, we just try again in a nanosecond.
	var cnt atomic.Int32

	cnt.Store(99)
	new := int32(123) // The value we're going to put into cnt via CAS

	for !cnt.CompareAndSwap(cnt.Load(), new) {
		time.Sleep(1 * time.Nanosecond)
	}

	fmt.Println(cnt.Load())
```

CAS data structures are tricky to implement correctly. See a CAS stack [here](./algorithms/col/cas/stack/stack.go). You simply **must** test CAS structures with tons of goroutines or else you won't spot real concurrency bugs and issues in your CAS code; otherwise, you'll definitely end up with some weird production bugs that don't make a lot of sense!

Some informal benchmark numbers of normal, RWMutex, and CAS for fun:
```
Push and pop 5000 elements, no goroutines
Normal   0.0000542 ns/op
RWMutex  0.0004570 ns/op
CAS      0.0001163 ns/op

Stringify 5000 elements, no goroutines
Normal   0.002163 ns/op
RWMutex  0.002093 ns/op
CAS      0.004539 ns/op

Push 5000 elements with 5000 goroutines, pop 5000 with no goroutines
RWMutex  0.002808 ns/op
CAS      0.001709 ns/op

Push 5000 elements with no goroutines, pop 5000 elements with 5000 goroutines
RWMutex  0.002587 ns/op
CAS      0.001841 ns/op

Normal == a normal stack with disabled RWMutex (skips locking)
RWMutex == a concurrent stack that uses RWMutex
CAS == a concurrent stack that uses CAS (no mutexes)
```

CAS is very approximately 30% faster for pushing and popping. Neat. CAS will definitely be slower for pushing many elements simultaneously or calling stringify since it's not a slice-backed data structure in my code.