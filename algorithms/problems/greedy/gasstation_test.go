package greedy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGasStation(t *testing.T) {
	gas := make([]int, 0, 5)
	dist := make([]int, 0, 5)

	gas = append(gas, 1, 2, 3, 4, 5)
	dist = append(dist, 3, 4, 5, 1, 2)

	actual := GasStation(gas, dist)

	assert.Equal(t, 3, actual)
}

// https://algo.monster/problems/gas_station
func GasStation(gas, dist []int) int {
	n := len(gas)

	for i := 0; i < n; i++ {
		start := i % n
		pos := start
		end := start + n
		fuel := 0

		fmt.Println("start", start, "pos", pos, "end", end)

		for ; pos < end; pos++ {
			fuel += gas[pos%n] - dist[pos%n]

			fmt.Println("\t", "gas", gas[pos%n], "dist", dist[pos%n], "fuel", fuel)

			if fuel < 0 {
				fmt.Println("\t Insufficient fuel, try next")
				break
			}
		}

		if pos == end {
			fmt.Println("Found the start", start)
			return start
		}
	}

	fmt.Println("No solution")
	return -1
}
