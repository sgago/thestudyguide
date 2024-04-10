package fibonaccinumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	What is Dynamic Programming?
	https://algo.monster/problems/dynamic_programming_intro

	The first Fibonacci numbers are:
		Fibonacci(0) = 0 // Base case
		Fibonacci(1) = 1 // Another base case
		Fibonacci(2) = 1 // First non base case
		Fibonacci(3) = 2 // 1 + 1
		Fibonacci(4) = 3 // 2 + 1
		Fibonacci(5) = 5 // 3 + 2
		Fibonacci(6) = 8 // 5 + 3
	or [0, 1, 1, 2, 3, 5, 8].

	We're programmers, so we start counting at 0...
*/

func TestFibonacciNumber2(t *testing.T) {
	num := 2
	actual := Fib(num) // 0+1

	fmt.Printf("Fib(%d)=%d\n", num, actual)
	assert.Equal(t, 1, actual)
}

func TestFibonacciNumber3(t *testing.T) {
	num := 3
	actual := Fib(num) // 0+1+1

	fmt.Printf("Fib(%d)=%d\n", num, actual)
	assert.Equal(t, 2, actual)
}

func TestFibonacciNumber5(t *testing.T) {
	num := 5
	actual := Fib(num)

	fmt.Printf("Fib(%d)=%d\n", num, actual)
	assert.Equal(t, 5, actual)
}

func TestFibonacciNumber6(t *testing.T) {
	num := 6
	actual := Fib(num)

	fmt.Printf("Fib(%d)=%d\n", num, actual)
	assert.Equal(t, 8, actual)
}

func TestFibonacciNumber10(t *testing.T) {
	num := 10
	actual := Fib(num)

	fmt.Printf("Fib(%d)=%d\n", num, actual)
	assert.Equal(t, 55, actual)
}

func Fib(n int) int {
	// Fib(0) and Fib(1) are equal to 0 and 1, respectively.
	// Return these base cases back immediately.
	if n <= 1 {
		return n
	}

	// Holds the solution to Fibonacci(n)
	// Initially Fibonacci(n-1) cause we haven't solved for Fibonacci(2) just yet...
	fibNum := 1

	memoOne := 1 // Holds the solution to Fibonacci(n-1)
	memoTwo := 1 // Holds the solution to Fibonacci(n-2)

	for i := 2; i < n; i++ {
		// Calculate our new Fibonacci number.
		fibNum = memoOne + memoTwo

		// Update the memo values
		// Pass memoOne to memoTwo and current fibNum to memoOne
		memoTwo, memoOne = memoOne, fibNum
	}

	return fibNum
}
