package containerwithmostwater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints
of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container,
such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
	Input: height = [1,8,6,2,5,4,8,3,7]
	Output: 49
	Explanation:
		The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
		In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
	Input: height = [1,1]
	Output: 1

Constraints:
- n == height.length
- 2 <= n <= 105
- 0 <= height[i] <= 104

==========================================================

Area will be calculated as min(y2, y1) * (x2 - x1). We want to max
this area equation. We need to loop through all the numbers.
We do not need to loop "backwards" because we'll have already tried it.
What do we do?
- We don't need to create all combos or anything like that.
- I don't really see an overlap in subproblems, per se. Like, why not two pointer?
- Going to do two pointers, I guess.

This is a moving in opposite direction, two pointer,
where we max out the area by keep the greatest height
column.
*/

func TestSolution(t *testing.T) {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	actual := maxArea(heights)

	fmt.Println("area:", actual)
	assert.Equal(t, 49, actual)
}

func maxArea(heights []int) int {
	ans := 0
	x1, x2 := 0, len(heights)-1

	for x1 < x2 {
		y1 := heights[x1]
		y2 := heights[x2]

		ans = max(ans, area(y2, y1, x2, x1))

		if y2 > y1 {
			x1++
		} else {
			x2--
		}
	}

	return ans
}

func area(y2, y1, x2, x1 int) int {
	return min(y2, y1) * (x2 - x1)
}
