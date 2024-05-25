package knapsackweightonly

import (
	"fmt"
	"sgago/thestudyguide/col/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Knapsack, Weight-Only
https://algo.monster/problems/knapsack_weight_only

Given a list of weights of n items,
find all sums that can be formed using their weights.

Input
	weights: A list of n positive integers,
	representing weights of the items
Output
	A list, in any order, of the unique sums that
	can be obtained by using combinations of the provided
	weights

Example 1:
	Input: weights = [1, 3, 3, 5]
	Output: [0, 1, 3, 4, 5, 6, 7, 8, 9, 11, 12]
	Explanation:
		We can form all sums from 0 to 12 except 2 and 10.
		Here is a short explanation for the sums:
		0: use none of the weights
		1: use item with weight 1
		3: use item with weight 3
		4: use weights 1 + 3 = 4
		5: use item with weight 5
		6: use weights 3 + 3 = 6
		7: use weights 1 + 3 + 3 = 7
		8: use weights 3 + 5 = 8
		9: use weights 1 + 3 + 5 = 9
		11: use weights 3 + 3 + 5 = 11
		12: use all weights

Constraints:
- 1 <= len(weights) <= 100
- 1 <= weights[i] <= 100
*/

/*
1. Consider []. We can make 0.
2. Consider [1]. We can make 0 and 1.
3. Consider [1 3]. We can make 0, 1, 3, and 4.

xxxxxxxx 0 1 2 3 4 5 6 7 8 9 a b c
NONE     T
1        T T
1 3      T T F T T
1 3 3    T T F T T F T T
1 3 3 5  T T F T T T T T T T F T T
*/
func TestKnapsackWeightOnly_Ugly_1335(t *testing.T) {
	actual := knapsackWeightOnly_Ugly([]int{1, 3, 3, 5})
	fmt.Println(actual)
}

func knapsackWeightOnly_Ugly(weights []int) []int {
	n := len(weights)
	totalSum := sum(weights)

	// row = weights
	// col = total
	dp := grid.New[bool](n+1, totalSum+1)

	// Init
	dp.Set(0, 0, true)

	fmt.Println(dp.String())

	// Each of the weights, by themselves, are "feasible"
	for w := 1; w < n+1; w++ {
		weight := weights[w-1]
		dp.Set(w, weight, true)
	}

	fmt.Println(dp.String())

	for w := 1; w < n+1; w++ {
		for j := 0; j < totalSum+1; j++ {
			prev := dp.Get(w-1, j)
			curr := dp.Get(w, j)
			dp.Set(w, j, curr || prev)
		}
	}

	fmt.Println(dp.String())

	// Ok, now let's do some adding
	for w := 1; w < n+1; w++ {
		curr := weights[w-1]

		for j := 0; j < totalSum+1; j++ {
			prev := dp.Get(w-1, j)

			if prev {
				dp.Set(w, curr+j, true)
			}
		}
	}

	// Now, let's carry thru prev solutions
	for w := 1; w < n+1; w++ {
		for j := 0; j < totalSum+1; j++ {
			prev := dp.Get(w-1, j)
			curr := dp.Get(w, j)
			dp.Set(w, j, curr || prev)
		}
	}

	fmt.Println(dp.String())

	ans := []int{}

	for i, val := range dp.LastRow() {
		if val {
			ans = append(ans, i)
		}
	}

	return ans
}

func TestKnapsackWeightOnly_1335(t *testing.T) {
	actual := knapsackWeightOnly([]int{1, 3, 3, 5})
	fmt.Println(actual)

	expected := []int{0, 1, 3, 4, 5, 6, 7, 8, 9, 11, 12}
	assert.Equal(t, expected, actual)
}

func knapsackWeightOnly(weights []int) []int {
	n := len(weights)
	totalSum := sum(weights)

	// row = weights
	// col = total
	// TODO: Optimize storage, can use bitmasking
	dp := grid.New[bool](n+1, totalSum+1)

	// Init
	dp.Set(0, 0, true)

	fmt.Println(dp.String())

	// Each of the weights, by themselves, are "feasible"
	for w := 1; w < n+1; w++ {
		curr := weights[w-1]
		dp.Set(w, curr, true)

		// Carry the previous solutions forward
		for j := 0; j < totalSum+1; j++ {
			prev := dp.Get(w-1, j)

			if prev {
				dp.Set(w, curr+j, true)
			}

			curr := dp.Get(w, j)
			dp.Set(w, j, curr || prev)
		}
	}

	fmt.Println(dp.String())

	ans := []int{}

	for i, val := range dp.LastRow() {
		if val {
			ans = append(ans, i)
		}
	}

	return ans
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
