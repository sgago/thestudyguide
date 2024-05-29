package numberofstepstoreduceanumberinbinaryrepresentationtoone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1404. Number of Steps to Reduce a Number in Binary Representation to One
https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one

Given the binary representation of an integer as a string s,
return the number of steps to reduce it to 1 under the following rules:
- If the current number is even, you have to divide it by 2.
- If the current number is odd, you have to add 1 to it.

It is guaranteed that you can always reach one for all test cases.

Example 1:
Input: s = "1101"
Output: 6
Explanation:
	"1101" corresponds to number 13 in their decimal representation.
	Step 1) 13 is odd, add 1 and obtain 14.
	Step 2) 14 is even, divide by 2 and obtain 7.
	Step 3) 7 is odd, add 1 and obtain 8.
	Step 4) 8 is even, divide by 2 and obtain 4.
	Step 5) 4 is even, divide by 2 and obtain 2.
	Step 6) 2 is even, divide by 2 and obtain 1.

Example 2:
Input: s = "10"
Output: 1
Explanation:
	"10" corresponds to number 2 in their decimal representation.
	Step 1) 2 is even, divide by 2 and obtain 1.

Example 3:
Input: s = "1"
Output: 0

Constraints:
- 1 <= s.length <= 500
- s consists of characters '0' or '1'
- s[0] == '1'
*/

/*
We're given a binary number like 1010.
If even (rightmost bit is 0), we have to divide.
If odd (rightmost bit is 1), we have to add one.
Return the number of steps to to reduce the number to 1.

Div by 2 is a bitshift right via >>, dropping the rightmost bit.
If odd, we add one.
Once we hit a power of two, it's turtles all the way down.

1101 = 13, +1
1110 = 14, >>
111 = 7, +1
1000 = 8, >>
100 = 4, >>
10 = 2, >>
1 = 1, >>

One soln is to convert the bin number to decimal and run the
operations until we get 1. n to get the decimal, then run like m ops. n+m
Can we do this in the bin string representation and get the time down to like n?
We can also run the bin ops on the string rep.
1
1
0 <- Flip all ones until we hit a zero
1 <- add one

What if we just track the number of odds in a counter?
Yeah, start at rightmost, if zero, delete it,

Until string is just "1"
- Start at rightmost
- If zero, delete it (a >> op)
- If one, store in counter
	- If next char is 1, it is actually a zero, delete it
	- If next char is 0, it is actually a one, flip it

We'll try this strategy.
*/

func TestZerolAndOneSub(t *testing.T) {
	fmt.Println('1' - '0') // 1
}

func TestNumberOfSteps_1101(t *testing.T) {
	actual := numberOfSteps("1101")
	fmt.Println(actual)

	assert.Equal(t, 6, actual)
}

func TestNumberOfSteps_1(t *testing.T) {
	actual := numberOfSteps("1")
	fmt.Println(actual)

	assert.Equal(t, 0, actual)
}

func numberOfSteps(b string) int {
	steps := 0
	odd := false

	for pos := len(b) - 1; pos > 0; {
		bit := b[pos]-'0' == 1

		if !bit && !odd || bit && odd {
			pos--
		} else {
			odd = !odd
		}

		steps++
	}

	return steps
}
