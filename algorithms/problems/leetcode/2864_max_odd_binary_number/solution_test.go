package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/*
2864. Maximum Odd Binary Number
You are given a binary string s that contains at least one '1'.
You have to rearrange the bits in such a way that the resulting binary
number is the maximum odd binary number that can be created from this combination.

Return a string representing the maximum odd binary number
that can be created from the given combination.

Note that the resulting string can have leading zeros.

Example 1:
	Input: s = "010"
	Output: "001"
	Explanation: Because there is just one '1',
	it must be in the last position. So the answer is "001".

Example 2:
	Input: s = "0101"
	Output: "1001"
	Explanation: One of the '1's must be in the last position.
	The maximum number that can be made with the remaining digits is "100". So the answer is "1001".


Constraints:
	1 <= s.length <= 100
	s consists only of '0' and '1'.
	s contains at least one '1'.

==========================================================================

Mmk, so odd binary strings always end with a 1.
1. Count number of 1's.
2. If it's "0100", then the answer is "0001"
3. If it's "0011", then the answer is "1001"
So, my string concat stuff would be like number_of_ones_minus_one + length_of_string_minus_number_of_ones + "1".
Ones are randomly arranged in the string, so we can just iterate through and count them.
I guess, we could convert to a number and try that, but it sounds like actual work.
So, loop thru and count ones. Then loop thru left side of string, converting 0's to 1's. Then add final one and return.
There's some loop iteration to sort out, but this feels pretty easy, so not going to implement it.
If string is length of 1, simply return the string.
Also, use strings.Builder to make this happen, to avoid frequent string concats.
*/

func TestMaxOddBinaryString(t *testing.T) {
	var sb strings.Builder

	// TODO: Take bin str and count ones... skipping cause boring lol
	// Is a 2N -> N solution with N space. Hmmm, we could probably, drop
	// the space consumption by sorting the string instead! Neat.
	inputStrLen := 20
	onesCount := 10

	for i := 0; i < onesCount-1; i++ {
		sb.WriteString("1")
	}

	for i := 0; i < inputStrLen-onesCount; i++ {
		sb.WriteString("0")
	}

	sb.WriteString("1")

	fmt.Println(sb.String())
}
