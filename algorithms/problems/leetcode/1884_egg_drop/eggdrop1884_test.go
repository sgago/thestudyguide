package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/*
	1884. Egg Drop With 2 Eggs and N Floors
	https://leetcode.com/problems/egg-drop-with-2-eggs-and-n-floors/description/

	1. We are given two identical eggs, whatever that means.
	2. We have access to a "building" with n floors labeled 1 to n.
	3. There exists a floor f where 0 <= f <= n such that an egg dropped
	at a floor higher than f will break and an egg at f floor or lower will not break.
	4. We may drop unbroken eggs from any floor x where 1 <=x <= n and if it's
	still unbroken we may reuse it.
	5. What's the minimum number of moves we can do to determine where f is.

	Mmk, so with one egg, we could simply drop the egg from each floor, start at
	1 drop it, goto 2, drop it, etc. If it breaks, then it's floor - 1. This is
	N search time.

	Now, they gave us two eggs. So, we can likely do some tricks to shorten the amount of guessing.

	Can we just do a binary search on this? There's an implicit boundary condition.

	At first:
	l
	      f
	0 1 2 3 4 5 6 7 8 9
                      r
			m

	Dropping the egg at 4, breaks it. Womp. One egg left. Updated pointers are:

	l
	      f
	0 1 2 3 4 5 6 7 8 9
          r
      m

	We'd do the remaining egg drop via bin search and would be good to go.

	Wait would, something starting like this would break both eggs:
	l
	f
	0 1 2 3 4 5 6 7 8 9
                      r
			m

	Then, one egg broke. So, next:
	l
	f
	0 1 2 3 4 5 6 7 8 9
          r
	  m

	Then, mmk, this works still.
	l
	f
	0 1 2 3 4 5 6 7 8 9
    r
	m

	I don't trust this, let's increase the size of our array:
	Drat floors start at 1...

	l
	f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
                                          r
					  m

	At a, the egg breaks, then
	l
	f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
                    r
	      m

	At 4, the egg breaks. We're out of eggs, so
	our remaining floors are 1 2 3. It looks like this:

	l
	f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
          r
	  m

	We have no way of continuing our algo. Womp.

	So, we can divide our floors with the first egg.
	Then we can find the exact floor with the remaining one.

	Say f is 99 always one less than the last floor:
	If we have 100 floors, and divide by 2, it's up to 50 guesses. 49+1.
	If we have 100 floors and divide by 10, it's up to 19 guesses. 10+9
	If we have 100 floors and divide by 5, it's up to 20+4 guesses.

	This looks like max is like n/x + d-1? 100/10 + 10-1 = 10+9 = 19?

	What's the min then of this f(x) = n/x + x - 1 where 1 <= x <= n?
	Or f(x) = 100/x + x - 1 where 1 <= x <= n?
	x = 100 -> 100
	x = 50 -> 51
	x = 25 -> 28
	x = 10 -> 19 <- Minima
	x = 5 -> 24

	100x^-1 + x - 1
	f'(x) = 1 - 100 / x^2.
	f'(x) = 0 = 1 - 100 / (10)^2

	Ahhh, but we can leep frog. We don't need to do 1 at a time and we don't need to check
	the last floor using leap frog.

	Our equation will be like sqrt(x)-1 + sqrt(x)/2.

	We can sqrt(n) to get optimal guesses? And then do the simple loop?

	Ugh, I misread the problem. The goal is actually not to find the actual floor
	f at all... it's to find the minimum number of moves to determine f lol.
	They even bolded that part of the problem! Ah, well, lessons re-learned then.
	Ah, also f can be 0. Must tread carefully and actually walk thru the problem.
	It's easy to miss this stuff, boundaries matter.

	Name our eggs x and y. x will go first, y after. Poor poor x...
	sqrt 20 = ~4.5 -> 5.

          f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
			x
	Egg breaks at 5, 1 guess.

	      f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
	y
	Egg doesn't break at 1, 2 guesses.

		  f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
	    y

	Egg doesn't break at 3, 3 guesses.
	Egg must be at 4. 5 and 3 didn't break. 3 guesses to find that out.

	      f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
			x

			      f
	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
	y   y

	If it breaks at 1, it's 0. Then, if it breaks at 3, it's 2... breaking is always f+1.
	sqrt(20)+sqrt(20)/2 ~= 4.5 + 4.5/2 ~= 6.75 ~= 7.
	sqrt(10)+sqrt(10)/2 ~= 3.2 + 3.2/2 ~= 4.8 ~= 5.

	Ah, but we don't have to do check the last of last floors.
	sqrt(20)-1+sqrt(20)/2 ~= 5.75 ~= 6. Man, you can use calculus to get this solution, instead of DP, pretty sure.

	90                  100
	o   x   x   x   x   x
	0 1 2 3 4 5 6 7 8 9 100
	9 interval guesses + 10/2...

	90                99
	o   x   x   x   x
	0 1 2 3 4 5 6 7 8 9
	9 interval guesses + 10/2-1

	If x is even... f(x) = (n/x-1) + x/2?
	If x is odd... f(x) = (n/x-1) + (x/2-1)?
*/

func TestTwoEggDropMath(t *testing.T) {
	n10 := twoEggDropMath(10)
	fmt.Println("10:", n10)

	n20 := twoEggDropMath(20)
	fmt.Println("20:", n20)

	n100 := twoEggDropMath(100)
	fmt.Println("100:", n100)

	n1000 := twoEggDropMath(1000)
	fmt.Println("1000:", n1000)
}

/*
      x     x     x
  1 2 3 4 5 6 7 8 9 10

        x       x       x       x       x
  1 2 3 4 5 6 7 8 9 a b c d e f g h i j k
									y

  9 + 5
  90...			 ...100
	  g   g   g   g   g
	1 2 3 4 5 6 7 8 9 100

	1 2 3 4 5 6 7 8 9 a b c d e f g h i j k 1 2 3 4 5 6 7 8 9 a b
*/

func twoEggDropMath(n int) int {
	// So, there's n floors and some optimal partitioning of floors,
	// let's call it x. The partitioning will then be n/x.
	// Because we need to check every partition no matter what,
	// n/x will be added to our total guesses.
	//
	// Then, within a partition, we can leap from over values.
	// You don't have to drop at every floor, you can skip a floor.
	// So, within the partition, we'll add x/2 more guesses.
	//
	// Our solution is then f(x) = n/x+x/2 = partition + leap frog.
	// Now, we want the lowest value. Deriving f(x) and solving
	// f'(x)=0, we can find a local minima.
	// f'(x)=1/2-n/x^2 and solving for x at 0, we get sqrt(2n).

	// Neat, but unfortunately, sqrt(2n) dumps a decimal value.
	// We can round that and get our solution.
	return int(math.Round(math.Sqrt(2 * float64(n))))
}

// Soln from interwebs that I think is off by 1, pretty sure:
func twoEggDropDp(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = i
	}

	dp[1][1] = 1

	for i := 2; i <= n; i++ {
		dp[i][1] = math.MaxInt
		for j := 1; j < i; j++ {
			curr := max(dp[j-1][0], dp[i-j][1]+1)
			dp[i][1] = min(dp[i][1], curr)
		}
	}

	return dp[n][1]
}

func TestTwoEggDropDp(t *testing.T) {
	n10 := twoEggDropDp(10)
	fmt.Println("10:", n10)

	n20 := twoEggDropDp(20)
	fmt.Println("20:", n20)

	// Good, this soln kicks out optimal as 13, but leetcode says it's 14.
	// That doesn't frustrate me at all... nope...
	n100 := twoEggDropDp(100)
	fmt.Println("100:", n100)

	n1000 := twoEggDropDp(1000)
	fmt.Println("1000:", n1000)
}
