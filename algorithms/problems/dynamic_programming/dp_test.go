package dynamicprogramming

import (
	"fmt"
	"sgago/thestudyguide/col/flags"
	"sgago/thestudyguide/col/grid"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
=================================
00. DYNAMIC PROGRAMMING INTRODUCTION
=================================
This is an in-depth guide for thinking through and solving dynamic programming (DP).
This guide assumes that you're decently comfortable with
- Using Go to run programs and/or tests
- 1D and 2D slices/arrays in Go
- Big 0 notation, time complexity (TC), and space complexity (SC)
- Binary searches
- Trees
- DFS, backtracking, and memoization

The name "dynamic programming" is 100% misnomer. There's nothing "dynamic" about it.
I would name it "fill up an array with answers to subproblems"
but that's a little long. Sarcasm aside, the rough idea really
is to fill up a slice such that an element has the final answer in it.

And, no, we're not going to start with annoying Fibonacci numbers that
an ancient Italian mathematician came up with in like 1200 AD.
I want to start with something simple. Dead. Simple.
Summing up values in a slice is what we're going to start with.
*/

/*
=================================
01. SUMMING A SLICE THE USUAL WAY
=================================
To begin easing into DP, let's sum up integers in an array the normal way.
*/
func Test_SumASlice_TheUsualWay(t *testing.T) {
	total := 0
	nums := []int{5, 3, 1, 4, 2} // Sums to 15

	for _, num := range nums {
		total += num
	}

	assert.Equal(t, 15, total) // Wow, exciting, amirite?
}

/*
=================================
02. SUMMING A SLICE THE DP WAY
=================================
Boring as summing up those numbers is, it is actually a basic
DP problem and solution.

The next test case demonstrates how to sum a slice "the DP way".

Instead of using a single integer, we're going to take a step in the
"wrong" direction and use a slice to solve the problem.
Each element will hold the sum of the previous numbers.
I always name the slice "dp". For this problem, the final element will hold our answer.

In our first couple of DP problems, we're going to use a 1D dp slice.
Later, we'll need to bust out 2D slices; for now, we're going to keep it simple.
Also, in our initial examples, the last element in slice will have the answer.
Again, this won't always be true, but, for now, it's a good starting point.

The problem wants us to find a sum, so our dp memo will also hold sums.
*/
func Test_SumASlice_TheDpWay(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2} // Sums to 15

	fmt.Println("nums:", nums, "sums to 15")

	// dp is our memo, that will hold subproblem sums.
	// dp[i] is going to hold the sum of all the previous numbers.
	dp := make([]int, len(nums))

	fmt.Println("  dp:", dp, "declare a memo called dp")

	// Our initial condition, dp[0], is going to be nums[0].
	dp[0] = nums[0]

	fmt.Println("  dp:", dp, "initial conditions! Store", nums[0], "at 0")

	// Our DP loop is going to start at idx 1 and
	// then look back at idx-1 to get the next sum.
	for i := 1; i < len(nums); i++ {
		prev := dp[i-1] // The previous sum from our memo
		curr := nums[i] // The next number to add

		dp[i] = curr + prev

		// Again, notice how dp[0], dp[1], dp[2], etc.
		// store answers to subproblems
		fmt.Println("  dp:", dp, curr, "+", prev, "store at index", i)
	}

	// And, as promised, the ans value holds our answer.
	ans := dp[len(dp)-1]
	fmt.Println("ans:", ans)
	assert.Equal(t, 15, ans)
}

// Same as above, but without distracting Printlns and comments.
func Test_SumASlice_TheDpWay_NoComments(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2} // Sums to 15

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + dp[i-1]
	}

	ans := dp[len(dp)-1]
	assert.Equal(t, 15, ans)
}

/*
=================================
03. WHY DID YOU SUM NUMBERS LIKE THAT!? IT LOOKS DUMB. WHY WOULD ANYONE DO THIS?
=================================
I know, I know. Hang in there. Again, DP is about filling up a slice with
answers to subproblems where one element has the final answer we're looking for.

Now, there's a couple of things you're going to want to get used to
doing *before* starting to actually code out a DP solution.

1. FIGURE OUT THE DP MEMO
For many problems, it's *usually* the same thing as whatever the problem is asking for.
- For summing a slice, we're asking for a sum, so our dp memo holds sums.
- If the problem is asking for true/false (feasibility), then dp memo will hold booleans.
- If the problem is asking for longest counts, then dp memo will hold integers of longest count of something
- If the problem is asking for a sum of coin values, then the dp memo usually holds coin value sums.

2. WRITE OUT THE DP MEMO STATES
We need to type out our DP memo - even the 2D ones - and any notes you need.
100% get use to doing this with the easy problems or you're
going to have a bad time when the problems get harder.
Doing this will get you comfortable with the problem statement nuances.

My memo notes for summing a slice containing 1 2 3 4 5 will look like this:
1 2 3 4  5  <- I typically write out the input slice because my memory is short
----------
0 0 0 0  0  <- Our initial dp memo of all zeros
1 0 0 0  0  <- 1 + 0 is one, so stick that in dp[0], this will actually be our initial condition
1 3 0 0  0  <- 1 + 2 is 3, so stick that solution in dp[1]
1 3 6 0  0  <- 3 + 3 = 6
1 3 6 10 0  <- 6 + 4 = 10
1 3 6 10 15 <- 10 + 5 = 15. Done, last element dp[4] has the answer of 15.

Ugly, but it's a good habit to get into. You'll thank me later.

3. DECLARE YOUR RECURRENCE RELATION
The entire goal of writing out the memo states is to a) get familiar
with the problem itself quickly and b) developing the recurrence
relation. The recurrence relation is a fancy name for a formula that gets your
dp memo from initial conditions to the next state, to the next state, and so on.
In other words, From dp[0] to dp[1] to dp[2] and so on.
To be blunt, this formula is critical. Once you have
and understand the recursion relation, you've got a big part of the problem solved.

Let's backtrack to summing up values in a slice.
For summing a slice, for each i-value, our recurrence relation is going to be
dp[0] = nums[1] <- nums[1] for our initial conditions, because our sum is zero initially
dp[1] = dp[1] + nums[2]
dp[2] = dp[2] + nums[3]
dp[3] = dp[3] + nums[4]
dp[4] = dp[4] + nums[5]

On occasion, you may find it helpful to declare the length of the dp memo
as n+1, where n is the length of the input slice. This is because we're going to
use the 0th index as a base condition. In memo notes, this would look similar to above:
dp[0] = 0
dp[1] = dp[0] + nums[1]
dp[2] = dp[1] + nums[2]
dp[3] = dp[2] + nums[3]
dp[4] = dp[3] + nums[4]
dp[5] = dp[4] + nums[5]

So, stated more generally, our recurrence relation is dp[n] = dp[n-1] + nums[n].
*/

/*
=================================
04. FIND THE MAX NUMBER
=================================
In this example, we show how you can get the max value in a slice "the DP way".
Just like the summing problem, we're going to use a dp memo to store the max value.
However, finding the max number introduces an important concept in DP: dealing with "overlap".
Many times when solving DP problems, we're going to need to look at multiple
previous dp values and figure out which one is the best to keep.

Say we have 3, 4, 5, 2, 1 and we want to find the maximum. As we iterate through the slice,
we need to make a "decision" of which number to keep at each step along the way.
Should we keep 3 or 4? 4 is bigger, so we keep 4. Should we keep 4 or 5? 5 is bigger, so we keep 5.

We will carry the max value all the way to the last element.

MEMO
Again, here's how the DP memo looks after each loop
3 0 0 0 0 <- Initial conditions, first value of 3.
3 4 0 0 0 <- Do we keep 3 or 4? 4 > 3, so keep 4
4 4 5 0 0 <- Do we keep 4 or 5? 5 > 4, so keep 5
4 4 5 5 0 <- Do we keep 5 or 2? 5 > 2, so keep 5
4 4 5 5 5 <- Do we keep 5 or 1? 5 > 1, so keep 5. Done.

RECURRENCE RELATION
Our recurrence relation is going to be dp[i] = max(dp[i-1], nums[i]).
*/
func Test_FindTheMaxNumber(t *testing.T) {
	nums := []int{3, 4, 5, 2, 1}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1])
	}

	last := dp[len(dp)-1]

	assert.Equal(t, 5, last)
}

// TRY IT: Find the minimum value in a slice using "the DP way", writing out the memo beforehand

/*
=================================
05. MINIMUM PATH SUM
=================================
This problem will introduce us to solving DP problems with 2D slices.
It will also combine the concepts of summing up values in a slice and finding the max value above.

Anyway, the minimum path sum problem is a classic DP problem.
The problem is to find the minimum path sum from the top left to the bottom right of a 2D slice.
So, we start in the upper left of a 2D slice and find the lowest cost to
get to the bottom right. You may only move right or down. And each move has the cost in the cell.

So, for
1 2 3
4 5 6
7 8 9

The minimum sum is 1 -> 2 -> 3 -> 6 -> 9 which sums to 21. This is the lowest cost.
*/
func Test_RobotPaths(t *testing.T) {
	path := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rows, cols := len(path), len(path[0])

	fmt.Println("Robot paths matrix being considered is")
	for r := 0; r < len(path); r++ {
		fmt.Println(path[r])
	}

	fmt.Println()

	// After you solve a bunch of 2D DP and graph problems with golang, you start to want a collection that can
	// - Declare a 2D slice with r rows and c cols instantly
	// - Allocs sub-slices immediately so that we don't have to think about associated errors
	// - Require the rows to be of equal length so that I don't have to think about those errors either
	// - Initializes all the elements to some value other than 0
	// - Lets you get any row or column as a single 1D slice
	// - Handles getting neighbor elements that are up, down, left, and right, either with row/col indexes or without
	// - Handles getting a sub-2D slice from a larger one, either with row/col indexes or without
	dp := grid.New[int](rows, cols)

	// Initial state at 0, 0
	dp.Set(0, 0, path[0][0])

	fmt.Println("Initial dp memo is")
	fmt.Println(dp.String())

	// The first row and column are annoying because
	// r-1 or c-1 will give us out-of-bound panics
	// (0 - 1 == -1 == invalid index)
	// We'll handle these with separate loops to keep the
	// cyclomatic complexity to a dull roar.

	// Fill in the first row
	for c := 1; c < cols; c++ {
		curr := path[0][c]
		prev := dp.Get(0, c-1)
		dp.Set(0, c, curr+prev)
	}

	fmt.Println("First row filled in")
	fmt.Println(dp.String())

	// Fill in the first column
	for r := 1; r < rows; r++ {
		prev := dp.Get(r-1, 0)
		curr := path[r][0]
		dp.Set(r, 0, curr+prev)
	}

	fmt.Println("First row and column filled in")
	fmt.Println(dp.String())

	// Now that we don't have to worry about
	// out-of-bound errors due to -1 indexes, fill in the rest
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			curr := path[r][c]

			prevLeft := dp.Get(r-1, c)
			prevUp := dp.Get(r, c-1)

			optimal := min(prevLeft, prevUp)

			dp.Set(r, c, optimal+curr)
		}

		fmt.Println("Row", r, "filled in")
		fmt.Println(dp.String())
	}

	// Answer is in the last cell value this time
	ans := dp.Last()
	fmt.Println("Last value, the answer:", ans)

	assert.Equal(t, 21, ans)
}

func Test_RobotPaths_NoComments(t *testing.T) {
	path := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rows, cols := len(path), len(path[0])

	dp := grid.New[int](rows, cols)

	// Initial state
	dp.Set(0, 0, path[0][0])

	// Fill in first row
	for c := 1; c < cols; c++ {
		dp.Set(0, c, path[0][c]+dp.Get(0, c-1))
	}

	// Fill in first column
	for r := 1; r < rows; r++ {
		dp.Set(r, 0, path[r][0]+dp.Get(r-1, 0))
	}

	// Fill in the rest
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			// Handle overlap, pick the smaller sum
			optimal := min(dp.Get(r-1, c), dp.Get(r, c-1))
			dp.Set(r, c, path[r][c]+optimal)
		}
	}

	ans := dp.Last()

	assert.Equal(t, 21, ans)
}

/*
=================================
COIN SUMS
=================================
Given unique integers (coins), find the number of
unique was we can sum the up to some total value.

For example, given coins with values 1, 2, 5 and a target total of 5, there
are 4 unique ways of summing 1, 2, and 5 to 5. They are:
- 1+1+1+1+1
- 2+1+1+1
- 2+2+1
- 5

The coin game piles on some new concepts:
 1. We loop through dp[0] to dp[n] multiple times, instead of just
    considering only dp[0] and then only dp[1] the dp[2] and so on.
    This helps gets us used to seeing and using our dp memo in a different way.
 3. Even the base conditions for the memo are tricky. We need to consider
    all the unique ways to make zero. There's 1 unique way to make zero and that's
    with no coins at all.

Our dp memo will hold the unique counts for each possible total.

The dp memo will look like:
[0 1 2 3 4 5] <- The sums we're trying to make, not the dp memo itself.
-------------
[1 0 0 0 0 0] <- Initial conditions, how many ways can we make zero. One way, with no coins at all.
[1 1 0 0 0 0] <- How many ways can we make 1 with the 1 value coin? 1 way, so increment dp[1].
[1 1 1 0 0 0] <- How many ways can we make 2 with the 1 value coin? 1 way, so increment dp[2].
[1 1 1 1 0 0] <- 3?
[1 1 1 1 1 0] <- 4?
[1 1 1 1 1 1] <- 5?
[1 1 2 1 1 1] <- How many ways can we make 2 with the two value coin? 1 way, so increment dp[1].
[1 1 2 2 1 1] <- 3?
[1 1 2 2 3 1] <- This part gets tricky. How many ways can we get 4? 1+1+1+1, 2+1+1, and 2+2! We don't just increment by 1 each time!
[1 1 2 2 3 3] <- Same for 5. We need to add 2!
[1 1 2 2 3 4]
[1 1 2 2 3 4]

Our recurrence relation is dp[i] = dp[i] + dp[i-coin] where coin <= i.
Stated generally, the formula is dp[i] += dp[i-coin].

Remember, we don't always just look at dp[4] once and move on.
Sometimes, we need to loop through the entire dp memo and maybe change each of the values.
*/
func Test_CoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5

	dp := make([]int, amount+1)
	dp[0] = 1

	fmt.Println("Initial memo is", dp)

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
			fmt.Println("Memo after considering subtotal", i, "with coin", coin, ":", dp)
		}

		fmt.Println("Memo after considering coin", coin, ":", dp)
	}

	ans := dp[len(dp)-1]

	fmt.Println("Total combinations:", ans)
	assert.Equal(t, 4, ans)
}

func Test_CoinChange_NoComments(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	ans := dp[len(dp)-1]
	assert.Equal(t, 4, ans)
}

/*
The longest increasing subsequence (LIS) problem piles on new DP concepts.
  - The final answer is not stored in the last element.
    We don't really know which number will be the last number in the
    longest increasing subsequence chain.
  - It has an outer an inner loop. Again, this is pretty common for DP problems.
    In this case, the outer loop moves the current number pointer forward.
    The inner loop walks backwards to find out the LIS is for the current number.
  - There is actually overlap to consider and deal with. For each of the numbers before the current one,
    if the previous number is smaller and it is the longest LIS we have, we'll record the value.

The LIS of [0, 1, 3, 2, 4, 5, -1, 0, 3] is 5. The longest subsequence is either
[0, 1, 2, 4, 5] or [0, 1, 3, 4, 5].

Our input slice is
[0, 1, 3, 2, 4, 5, -1, 0, 3]

The DP memo looks like this after each loop.
The _ underscores are just for visual separation:

1 1 1 _ 1 1 1 _ 1 1 1  <- Our initial dp memo
1 2 1 _ 1 1 1 _ 1 1 1  <- 0 to 1 is a LIS of 2.
1 2 3 _ 1 1 1 _ 1 1 1  <- 0 to 1 to 3 is a LIS of 3.
1 2 3 _ 3 1 1 _ 1 1 1  <- 0 to 1 to 2 is a LIS of 3.
1 2 3 _ 3 4 1 _ 1 1 1
1 2 3 _ 3 4 5 _ 1 1 1  <- We're done here... we just don't "know" that yet
1 2 3 _ 3 4 5 _ 1 1 1
1 2 3 _ 3 4 5 _ 1 2 1
1 2 3 _ 3 4 5 _ 1 2 4  <- Done. 5 is biggest LIS count with values 0 1 2 4 5 or 0 1 3 4 5

The recurrence relation is like dp[i] = max(dp[i], dp[prev_i] + 1) but only dp[prev_i] where
nums[i] > nums[prev_i], because otherwise it wouldn't be only increasing numbers.
*/
func Test_LongestIncreasingSubsequence(t *testing.T) {
	nums := []int{0, 1, 3, 2, 4, 5, -1, 0, 3} // LIS == 5 == 0, 1, 3, 4, 5

	// Our memo dp again. It's going to store the LIS for every single element
	// in the slice. But the final index won't have our answer because we don't
	// know which number is actually the final number in the longest increasing chain.
	dp := make([]int, len(nums))

	// Each number is minimally an LIS of 1 all by itself, so
	// for our initial conditions, we'll just set everything to 1.
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		// INNER LOOP
		// An inner loop is common for most DP problems.
		// This one walks backwards to find the the LIS for the current number.
		for j := i - 1; j >= 0; j-- {
			prev := nums[j]

			if curr > prev {
				// So, only if the current number is bigger than the previous
				// one, will we consider looking at previous dp[j] values.
				// Otherwise, if curr <= prev, it wouldn't be an LIS; it would be equal
				// to or decreasing instead!

				// Now, we add +1 to the previous LIS value because we're adding
				// curr to the LIS chain.

				// But dp[j] + 1 might not be the longest! We need to walk back over
				// all previous dp values to find the LIS.
				maybeLonger := dp[j] + 1

				// OVERLAP
				// We need curr > prev, but we also have some overlap to deal with
				// Is the current value in dp[i] longer? It might be! Or is it
				// dp[j]+1 longer? We'll keep which ever number is bigger because
				// we want to find the LIS.
				dp[i] = max(dp[i], maybeLonger)
			}
		}
	}

	fmt.Println("nums:", nums)
	fmt.Println("  dp:", dp)

	// ANSWER
	// The answer might not be the last number for this one!
	// We'll loop thru our dp memo and find the answer.
	longest := slices.Max(dp)

	fmt.Println(" ans:", longest)
	assert.Equal(t, 5, longest)
}

func Test_LongestIncreasingSubsequence_NoComments(t *testing.T) {
	nums := []int{0, 1, 3, 2, 4, 5, -1, 0, 3} // LIS == 5 == 0, 1, 3, 4, 5

	dp := make([]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		for j := i - 1; j >= 0; j-- {
			prev := nums[j]

			if curr > prev {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	longest := slices.Max(dp)

	assert.Equal(t, 5, longest)
}

/*
Perfect squares DP problem is going to count the number of
squared numbers we need to use to sun up some other number.

In classic competitive programming style, we get some
out-of-left-field feeling math concepts that some PhD bothered to give a name to.
Like the "perfect square" and need to understand it quickly.

There's nothing perfect about the square at all, really. It's literally
2 * 2 = 4, 3 * 3 = 9, 4 * 4 = 16, 5 * 5 = 25, etc. That's all they are,
a number multiplied by itself aka n^2 is a "perfect square".

So, perfect square values themselves are 1, 4, 9, 16, 25, 36, etc.
And, we're going to count the minimum number of perfect squares used to get
some other number.

Examples of numbers and counts of perfect squares used to compute them:
- 7 = 4 + 1 + 1 + 1 = 4 squares used
- 9 = 9 = 1 square used
- 10 = 9 + 1 = 2 squares used
- 12 = 4 + 4 + 4 = 3 squares used
- 13 = 9 + 4 = 2 squares used
- 15 = 9 + 4 + 1 + 1 = 4 squares used

This problem will use similar DP concepts from LIS.
But it also adds creating both the input slice instead of being given one,
unlike LIS. Also, the DP memo values is a bit less obvious.

Our memo will look something like below, where x is just the previous value from above.
I'm trying to show how the dp memo evolves here without writing every
state cause there would be a ton of states to show (~n*sqrt(n) ish?).

Again, x is a previous value and _ underscores are for visual separation.

0 1 2 3 _ 4 5 6 _ 7 8 9 _ 10 11 12 _ 13 14 15  <- Just using 1 = 1 * 1
0 x x x _ 1 2 3 _ 4 2 3 _  4  5  3 _  4  5  6  <- Better counts using 4 = 2 * 2, we get some overlap here, using 4 for 4 is better than 1 + 1 + 1 + 1.
0 x x x _ x x x _ x x 1 _  2  3  4 _  2  3  4  <- Even better counts using 9 = 3 * 3
0 1 2 3 _ 1 2 3 _ 4 2 1 _  2  3  4 _  2  3  4  <- Final DP memo, after we pick the min value for each case
*/
func Test_PerfectSquares(t *testing.T) {
	num := 15

	// Init our memo with the maximum squares
	// required for any ith value.
	// The maximum squares for any value is using all 1s
	// so we initialize our dp memo with 1 * i == i.
	dp := make([]int, num+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}

	iter := 0

	for i := 2; i*i <= num; i++ {
		sqr := i * i

		for j := sqr; j <= num; j++ {
			iter++

			prev := dp[j]
			maybeLess := dp[j-sqr] + 1

			// Here, is our recurrence relation and our way of addressing any overlap.
			// Can we get to dp[j] in some more efficient way?
			// For example, from our initial dp slices above, can we compute
			// 4 using less than 4 1s? Yes, we can use 2*2 = 4, so only 1 square (2*2)
			// is required. Then, can we get to 5, using less than 5 1s? Yes, 4+1= 2*2+1*1.
			dp[j] = min(prev, maybeLess)

			fmt.Println("dp:", dp)
		}
	}

	ans := dp[len(dp)-1]

	fmt.Println("last:", ans)
	fmt.Println("iter:", iter)
}

/*
THE DIVISOR GAME

The divisor game is a special. You see, unlike other games you may have played,
this one is not fun to play with friends lol.

Anyway, each player will divide a number n by another number x
such that n/x does not result in a remainder (divides evenly), where 1 < x < n. Note that
n may not equal x. Then, we subtract x from n and the next
player takes their turn with the remaining n value. When the player can no longer divide x by anything, when n == 1,
that player loses.

Let's go through a bunch of cases to get a feel for the game:
  - If player 1 gets an n == 1, they lose automatically. There is no number smaller than 1 that we can pick.
  - If player 1 gets an n == 2, they win. They choose 1, subtract 2 - 1 = 1, the other player gets 1 and loses.
  - If player 1 gets an n == 3, they lose. It's 3/1 = 2, player 2 does 2/1, and player 1 gets 1 and loses.
  - If player 1 gets an n == 4, they finally have a choice. They can pick 1 or 2. 1 gives the other player a 3 so they lose.
    If player 1 would pick 2, they would give player 2 a 2 and then player 1 would lose.
  - If player 1 gets an n == 5, this is prime and they don't have a choice. Player 1 has to pick 1 and gives
    player 2 a 4. From our prior work (n==4 above), we know player 2 would choose 1 to win.
  - If player 1 gets an n == 6, we can pick 1, 2, and 3 which would give player 2 a 5, 4, or 3, respectively.
    Player 1 can choose 3, give player 2 a 3, they must choose 1, player 1 gets a 2, player 1 chooses 1, and player 2 loses
    with 1 remaining.

From running these through, at any point player 1 can get a win from a subcase, they just win.

Our recurrence relation is something like dp(n) == dp(n-1) || dp(n-2) || dp(1) but only take terms
where n%dp(n-1)==0, n%dp(n-2)==0, etc.

Our initial conditions can be whether player 1 wins, dp(1) = f, dp(2) = t, dp(3) = f.
We can do !dp(n) to determine if player 2 wins for a particular state. Player 1 does NOT want
player 2 to win, so we only wanna pick numbers where player 1 wins.
*/
func Test_TheDivisorGame(t *testing.T) {
	n := 8

	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = false
	dp[2] = true
	dp[3] = false

	for i := 4; i <= n; i++ {

		for j := 1; j < i; j++ {

			noRemainder := i%j == 0
			player2DoesntWin := !dp[i-j]

			if noRemainder && player2DoesntWin {
				dp[i] = true

				// Player 1 can force a win by choosing
				// this ith value, so just mark true and quit
				// (Yes, we can move this break into the for-loop condition)
				break
			}
		}
	}

	fmt.Println("Player 1 can force a win:", dp[len(dp)-1])
}

func Test_TheDivisorGame_MemoryOptimizations(t *testing.T) {
	n := 8

	// So, most of the time, if you see a slice of bools like
	// flags := make([]bool, 10), it can be replaced with
	// bit flagging. Each bool takes 1 byte, and typically malloc() won't give you
	// memory with weird boundaries like 1 or 3 bits.
	// Your actual choices are always like 8, 16, 32, 64, and maybe 128. So, a bool is 1 byte.
	// Therefore, we can save jam a bunch of bit values into an unsigned integer.
	// This is more space efficient but more annoying to work with.

	// To make it even easier, I've implemented a bit flags struct already.
	// It'll pick uint 8, 16, 32, or 64 for you and append more unsigned integers as needed.
	dp := flags.New(n + 1)
	dp.Set(0, false)
	dp.Set(1, false)
	dp.Set(2, true)
	dp.Set(3, false)

	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			noRemainder := i%j == 0
			player2DoesntWin := !dp.Get(i - j)

			if noRemainder && player2DoesntWin {
				dp.Set(i, true)
				break
			}
		}
	}

	fmt.Println("Player 1 can force a win:", dp.Get(n))
}

func Test_LongestCommonSubsequence(t *testing.T) {
	w1 := "aabcae"
	w2 := "ace"

	nw1 := len(w1)
	nw2 := len(w2)

	dp := grid.
		New[int](nw1+1, nw2+1)

	fmt.Println(dp.String())

	for i := 1; i < nw1+1; i++ {
		r1 := w1[i-1]

		for j := 1; j < nw2+1; j++ {
			r2 := w2[j-1]

			if r1 == r2 {
				prev := dp.Get(i-1, j-1)
				dp.Set(i, j, prev+1)
			} else {
				left := dp.Get(i, j-1)
				up := dp.Get(i-1, j)

				dp.Set(i, j, max(left, up))
			}

			fmt.Println(dp.String())
		}
	}

	ans := dp.Last()

	fmt.Println("ans:", ans)
	assert.Equal(t, 3, ans)
}

func Test_LongestCommonSubsequence_WithMemoryOptimization(t *testing.T) {
	word1 := "aabcae"
	word2 := "ace"

	n := len(word1)
	m := len(word2)

	dp := grid.
		New[int](2, m+1)

	fmt.Println(dp.String())

	nextI := 0

	for i := 1; i < n+1; i++ {
		// Using these vars for indexes,
		// we can just ping-pong between rows
		// 0 and 1 without moving data or creating new slices
		prevI := 0
		nextI = i % 2

		if nextI == 0 {
			prevI = 1
		}

		r1 := word1[i-1]

		for j := 1; j < m+1; j++ {
			r2 := word2[j-1]

			if r1 == r2 {
				prev := dp.Get(prevI, j-1)
				dp.Set(nextI, j, prev+1)
			} else {
				left := dp.Get(nextI, j-1)
				vert := dp.Get(prevI, j)

				dp.Set(nextI, j, max(left, vert))
			}

			fmt.Println(dp.String())
		}
	}

	ans := dp.Row(nextI)[dp.Cols()-1]

	fmt.Println("ans:", ans)
	assert.Equal(t, 3, ans)
}
