// Package lamport provides a simple implementation of the Lamport logical clock.
package lamport

import (
	"strconv"
	"sync/atomic"
)

// Clock represents a Lamport logical clock.
type Clock struct {
	ticks atomic.Uint64
}

// New creates a new instance of the Clock.
func New() *Clock {
	return &Clock{}
}

// Inc increments the clock by 1 and returns the new value.
func (c *Clock) Inc() uint64 {
	return c.ticks.Add(1)
}

// Set sets the clock to the specified value.
func (c *Clock) Set(ticks uint64) {
	old := c.ticks.Load()

	for !c.ticks.CompareAndSwap(old, ticks) {
		old = c.ticks.Load()
	}
}

// Get returns the current value of the clock.
func (c *Clock) Get() uint64 {
	return c.ticks.Load()
}

// Reset resets the clock to 0.
func (c *Clock) Reset() {
	c.ticks.Store(0)
}

func (c *Clock) String() string {
	return strconv.FormatUint(c.ticks.Load(), 10)
}
