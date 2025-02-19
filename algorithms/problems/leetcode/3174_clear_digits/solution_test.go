package cleardigits

import (
	"fmt"
	"testing"
)

/*
3174. Clear Digits
https://leetcode.com/problems/clear-digits

You are given a string s.

Your task is to remove all digits by doing this operation repeatedly:
- Delete the first digit and the closest non-digit character to its left.

Return the resulting string after removing all digits.

Example 1:
	Input: s = "abc"
	Output: "abc"
	Explanation:
	There is no digit in the string.

Example 2:
	Input: s = "cb34"
	Output: ""
	Explanation:
	First, we apply the operation on s[2], and s becomes "c4".
	Then we apply the operation on s[1], and s becomes "".

Constraints:
- 1 <= s.length <= 100
- s consists only of lowercase English letters and digits.
- The input is generated such that it is possible to delete all digits.

========================================================================

So, from the left side of the string, we need to remove the first digit,
and the closest non-digit character to its left.
There's a few ways we can solve though it looks like a two-pointer problem.
A slow + fast pointer. Slow will mark a non-digit, fast will find the next digit.
The slow one won't stay at the beginning, it can be moved along with fast to the next non-digit.
Another option is to use a stack to keep track of the letters and compare the last letter of part.
*/

func TestClearDigitsInMemory(t *testing.T) {
	s := "cb34"

	actual := ClearDigitsInMemory(s)
	fmt.Println(actual)
}

func BenchmarkClearDigitsInMemory(b *testing.B) {
	s := "cb34"
	for i := 0; i < b.N; i++ {
		_ = ClearDigitsInMemory(s)
	}
}

func ClearDigitsInMemory(s string) string {
	slow, fast := 0, 0

	for ; fast < len(s); fast++ {
		if slow != fast && s[fast] >= '0' && s[fast] <= '9' {
			fast--
			s = s[:slow] + s[slow+1:]
			s = s[:fast] + s[fast+1:]
		} else {
			slow = fast
		}
	}

	return s
}

func TestClearDigitsStack(t *testing.T) {
	s := "cb34"

	actual := ClearDigitsInMemory(s)
	fmt.Println(actual)
}

func BenchmarkClearDigitsStack(b *testing.B) {
	s := "cb34"
	for i := 0; i < b.N; i++ {
		_ = ClearDigitsStack(s)
	}
}

func ClearDigitsStack(s string) string {
	stack := make([]rune, 0, len(s))

	for _, r := range s {
		if r >= '0' && r <= '9' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	return string(stack)
}

/*
BenchmarkClearDigitsInMemory-10    	27008818	        42.90 ns/op	       6 B/op	       2 allocs/op
BenchmarkClearDigitsStack-10       	56808766	        20.73 ns/op	      16 B/op	       1 allocs/op

Neat, the stack solution is faster and uses less memory.
*/
