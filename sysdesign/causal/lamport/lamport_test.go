package lamport

import (
	"sync"
	"testing"
)

func TestSetConcurrent(t *testing.T) {
	clock := New()

	// Set the clock to an initial value
	clock.Set(0)

	// Number of goroutines to spawn
	numGoroutines := 1000

	// Wait group to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrently set the clock to different values
	for i := 0; i < numGoroutines; i++ {
		go func(value uint64) {
			clock.Set(value)
			wg.Done()
		}(uint64(i))
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify that the clock value is set correctly
	if clock.Get() != uint64(numGoroutines-1) {
		t.Errorf("Expected clock value to be %d, but got %d", numGoroutines-1, clock.Get())
	}
}
