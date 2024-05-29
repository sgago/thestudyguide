package studentattendancerecordii

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
552. Student Attendance Record II
https://leetcode.com/problems/student-attendance-record-ii

An attendance record for a student can be represented as a
string where each character signifies whether the student
was absent, late, or present on that day. The record only
contains the following three characters:

'A': Absent.
'L': Late.
'P': Present.

Any student is eligible for an attendance award if they
meet both of the following criteria:

The student was absent ('A') for strictly fewer than 2 days total.
The student was never late ('L') for 3 or more consecutive days.
Given an integer n, return the number of possible attendance
records of length n that make a student eligible for an
attendance award. The answer may be very large, so return i
modulo 109 + 7.

Example 1:
Input: n = 2
Output: 8
Explanation: There are 8 records with length 2 that are eligible
for an award:
"PP", "AP", "PA", "LP", "PL", "AL", "LA", "LL"
Only "AA" is not eligible because there are 2
absences (there need to be fewer than 2).

Example 2:
Input: n = 1
Output: 3

Example 3:
Input: n = 10101
Output: 183236316

Constraints:
- 1 <= n <= 10^5
*/

/*
Learnings:
- Modulus rule: (a+b)%m = (a%m+b%m)%m
	You can spread modulus to the a and b, as long as you mod the addition result
- This one was actually much easier with recursion than looping.
*/

func TestStudentAttendence2_1day(t *testing.T) {
	actual := studentAttendance2(1)

	fmt.Println(actual)
	assert.Equal(t, 3, actual)
}

func TestStudentAttendence2_2days(t *testing.T) {
	actual := studentAttendance2(2)

	fmt.Println(actual)
	assert.Equal(t, 8, actual)
}

func TestStudentAttendence2_10101days(t *testing.T) {
	actual := studentAttendance2(10101)

	fmt.Println(actual)
	assert.Equal(t, 183236316, actual)
}

// TODO: Shrink state used for memo, we can do this
// with a single int, pretty sure, like
// (i*100)+(A*10)+L
type state struct {
	i int
	A int
	L int
}

func studentAttendance2(days int) int {
	mod := int(math.Pow10(9)) + 7
	var inner func(i, A, L int) int
	memo := make(map[state]int)

	inner = func(i int, A int, L int) int {
		if cnt, ok := memo[state{i, A, L}]; ok {
			return cnt
		}

		if A >= 2 {
			return 0
		} else if L >= 3 {
			return 0
		} else if days == i {
			return 1
		}

		// Add P
		ans := inner(i+1, A, 0) % mod

		// Add A
		ans += inner(i+1, A+1, 0) % mod

		// Add L
		ans += inner(i+1, A, L+1) % mod

		memo[state{i, A, L}] = ans % mod

		return ans % mod
	}

	return inner(0, 0, 0)
}
