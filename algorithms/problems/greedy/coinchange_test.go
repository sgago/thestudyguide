package greedy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreedyCoinChange(t *testing.T) {
	coins := []int{100, 50, 10, 5, 1}
	target := 369

	count := GreedyCoinChange(target, coins...)

	assert.Equal(t, 10, count)
}

func TestGreedyCoinChange_ShowFailure(t *testing.T) {
	coins := []int{30, 20, 1}
	target := 40

	count := GreedyCoinChange(target, coins...)

	assert.Equal(t, 11, count)
}

func GreedyCoinChange(target int, coins ...int) int {
	count := 0

	for _, coin := range coins {
		for target-coin >= 0 {
			target -= coin
			count++
		}
	}

	return count
}

func DpCoinChange(target int, coins ...int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // How many ways to make 0? 1 way.

	for _, coin := range coins {
		for i := coin; i <= target; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[target]
}

func BruteCoinChange(target int, coins ...int) int {
	count := math.MaxInt

	for i := 1; i < len(coins); i++ {
		sum := target
		cnt := 0

		for j := i; j > 0; j-- {
			coin := coins[j]

			cnt += sum / coin
			sum := sum % coin

			if sum == 0 {
				break
			}
		}

		count = min(count, cnt)
	}

	return count
}

func TestDpCoinChange(t *testing.T) {
	coins := []int{1, 20, 30}
	target := 40

	count := DpCoinChange(target, coins...)

	assert.Equal(t, 2, count)
}

func TestBruteCoinChange_ShowSuccess(t *testing.T) {
	coins := []int{1, 20, 30}
	target := 40

	count := BruteCoinChange(target, coins...)

	assert.Equal(t, 2, count)
}

// Benchmark for DpCoinChange
func BenchmarkDpCoinChange(b *testing.B) {
	target := 100
	coins := []int{1, 2, 5, 10, 25, 50}

	for i := 0; i < b.N; i++ {
		DpCoinChange(target, coins...)
	}
}

// Benchmark for BruteCoinChange
func BenchmarkBruteCoinChange(b *testing.B) {
	target := 100
	coins := []int{1, 2, 5, 10, 25, 50}

	for i := 0; i < b.N; i++ {
		BruteCoinChange(target, coins...)
	}
}

// Benchmark for DpCoinChange with large target and coins set
func BenchmarkDpCoinChange_Large(b *testing.B) {
	target := 10000
	coins := []int{1, 5, 10, 25, 50, 100, 200, 500, 1000}

	for i := 0; i < b.N; i++ {
		DpCoinChange(target, coins...)
	}
}

// Benchmark for BruteCoinChange with large target and coins set
func BenchmarkBruteCoinChange_Large(b *testing.B) {
	target := 10000
	coins := []int{1, 5, 10, 25, 50, 100, 200, 500, 1000}

	for i := 0; i < b.N; i++ {
		BruteCoinChange(target, coins...)
	}
}

// Benchmark for DpCoinChange with very large target and more coins
func BenchmarkDpCoinChange_VeryLarge(b *testing.B) {
	target := 50000
	coins := []int{1, 2, 5, 10, 25, 50, 100, 200, 500, 1000, 2000, 5000}

	for i := 0; i < b.N; i++ {
		DpCoinChange(target, coins...)
	}
}

// Benchmark for BruteCoinChange with very large target and more coins
func BenchmarkBruteCoinChange_VeryLarge(b *testing.B) {
	target := 50000
	coins := []int{1, 2, 5, 10, 25, 50, 100, 200, 500, 1000, 2000, 5000}

	for i := 0; i < b.N; i++ {
		BruteCoinChange(target, coins...)
	}
}
