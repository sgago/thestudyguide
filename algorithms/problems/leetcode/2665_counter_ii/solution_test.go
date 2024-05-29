package counterii

import (
	"fmt"
	"testing"
)

/*
2665. Counter II
https://leetcode.com/problems/counter-ii

Write a function createCounter. It should accept an initial integer init.
It should return an object with three functions.

The three functions are:
increment() increases the current value by 1 and then returns it.
decrement() reduces the current value by 1 and then returns it.
reset() sets the current value to init and then returns it.

Example 1:
	Input: init = 5, calls = ["increment","reset","decrement"]
	Output: [6,5,4]
	Explanation:
		const counter = createCounter(5);
		counter.increment(); // 6
		counter.reset(); // 5
		counter.decrement(); // 4

Example 2:
	Input: init = 0, calls = ["increment","increment","decrement","reset","reset"]
	Output: [1,2,1,0,0]
	Explanation:
		const counter = createCounter(0);
		counter.increment(); // 1
		counter.increment(); // 2
		counter.decrement(); // 1
		counter.reset(); // 0
		counter.reset(); // 0

Constraints:
- -1000 <= init <= 1000
- 0 <= calls.length <= 1000
- calls[i] is one of "increment", "decrement" and "reset"
*/

/*
Huh, weird problem. Looks like it's just a JS/TS problem maybe?
Well, whatever. Here we go!
*/

func TestCounterII(t *testing.T) {
	runCounterII(5, []string{"increment", "reset", "decrement"})
}

func runCounterII(init int, calls []string) {
	c := New(init)

	for i, call := range calls {
		var val int

		switch call {
		case "increment":
			val = c.Inc()
		case "decrement":
			val = c.Dec()
		case "reset":
			val = c.Reset()
		default:
			panic(fmt.Sprintf("call %s is invalid", call))
		}

		fmt.Println("i:", i, "call:", call, "counter:", val)
	}
}

type counter struct {
	init int
	val  int
}

func New(init int) *counter {
	return &counter{
		init: init,
		val:  init,
	}
}

func (c *counter) Inc() int {
	c.val++
	return c.val
}

func (c *counter) Dec() int {
	c.val--
	return c.val
}

func (c *counter) Reset() int {
	c.val = c.init
	return c.val
}
