package lock

import "sync"

// RW is a reader/writer mutual exclusion locker for concurrent access.
type RW struct {
	// Indicates if this locker is enabled for synchronizing access.
	sync bool

	// Mutex used for synchronizing.
	mu sync.RWMutex
}

// Read returns a function that acquires a read lock.
// Example:
//
//	defer lock.Read()()
func (lock *RW) Read() func() {
	if !lock.sync {
		return func() {}
	}

	lock.mu.RLock()
	return lock.mu.RUnlock
}

// Write returns a function that acquires a write lock.
// Example:
//
//	defer lock.Write()()
func (lock *RW) Write() func() {
	if !lock.sync {
		return func() {}
	}

	lock.mu.Lock()
	return lock.mu.Unlock
}

// Synchronize sets the synchronization flag for this locker.
// If enable is true, it enables synchronization, meaning that
// the locker will synchronize concurrent access using locks.
// If enable is false, it disables synchronization, allowing
// concurrent access without additional synchronization.
func (lock *RW) Synchronize(enable bool) {
	lock.sync = enable
}

// IsSynchronized returns true if the locker is configured for synchronization,
// meaning that concurrent access is synchronized using locks.
// Returns false if synchronization is disabled,
// allowing concurrent access without additional synchronization.
func (lock *RW) IsSynchronized() bool {
	return lock.sync
}
