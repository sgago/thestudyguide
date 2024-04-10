package leetcode

import (
	"fmt"
	"strconv"
	"testing"

	"sgago/thestudyguide/col/queue"
)

/*
	Operators: +, -, *, "".
	DFS + permutation it. Iterative dev: add, multi, multi-digit numbers, subtraction.
	State: Stringified expression, agg sum of everything.
*/

func TestExpressionAddOperator(t *testing.T) {
	num := "234"
	target := 8

	actual := expressionAddOperator(num, target)

	fmt.Printf("%v\n", actual)
}

// TODO: Make ops consts
// TODO: Reorg file
var operators = []string{"+", "*", ""}

type state struct {
	index      int
	total      int
	prev       int
	prevOp     string
	expression string
}

func expressionAddOperator(num string, target int) []string {
	digits := len(num)
	results := make([]string, 0)
	all := make([]string, len(num)*len(operators))

	// TODO: cleanup rune to string to int convert, it's gross
	first, _ := strconv.Atoi(string(num[0]))

	q := queue.New[state](1, state{
		index:      0,
		total:      first,
		prev:       0,
		prevOp:     "",
		expression: string(num[0]),
	})

	for !q.Empty() {
		curr := q.DeqHead()

		if curr.index < digits-1 {
			next, _ := strconv.Atoi(string(num[curr.index+1]))

			// TODO: Reduce code dupe
			for _, op := range operators {
				switch op {
				case "+":
					q.EnqHead(state{
						index:      curr.index + 1,
						total:      curr.total + next,
						prev:       curr.total,
						prevOp:     "+",
						expression: curr.expression + "+" + string(num[curr.index+1]),
					})
				case "*":
					// 2 + 3 * 2
					// 5 * 2 WRONG, 3 * 2 RIGHT
					// Total = 5
					// Prev = 2a
					// Next = also 2b
					// (total - prev) * next + prev?
					// (5 - 2a) * 2b + 2a? = 8?
					//
					q.EnqHead(state{
						index:      curr.index + 1,
						total:      (curr.total-curr.prev)*next + curr.prev,
						prev:       0,
						prevOp:     "*",
						expression: curr.expression + "*" + string(num[curr.index+1]),
					})
				case "":
					var newTotal int
					if curr.prevOp == "+" {
						newTotal = (curr.total-curr.prev)*10 + curr.prev + next
					} else if curr.prevOp == "*" && curr.prev != 0 {
						newTotal = (curr.total/curr.prev)*10 + curr.prev + next
					} else if curr.prevOp == "" {
						newTotal = (curr.total)*10 + curr.prev + next
					}

					// Total 2
					// Prev 0
					// Next is 3
					// Is this actual the same thing again? Is it:
					// (total - prev) + prev * 10 + next?
					// Yeah, this is just multi again, but always *10...
					// Ugh, what do I do about previous?
					// 2*34=64 WRONG
					// 2 = 2
					// 2 * 3 = 6
					// 2 * 34 = Back out previous op 6 = 2 * 3
					// Yeah, we can't just always subtract, we need to divide too
					// (6 / 2) * 10 + 4
					// For multi:
					// total = (total / prev) * 10 + next
					// For add:
					// total = (total - prev) * 10 + next
					// Track previous operator
					q.EnqHead(state{
						index:      curr.index + 1,
						total:      newTotal,
						prev:       curr.prev,
						prevOp:     "",
						expression: curr.expression + string(num[curr.index+1]),
					})
				}
			}
		} else if curr.total == target {
			results = append(results, curr.expression)
		}

		all = append(all, fmt.Sprintf("%s=%d", curr.expression, curr.total))
	}

	for _, e := range all {
		fmt.Printf("%v\n", e)
	}

	return results
}
