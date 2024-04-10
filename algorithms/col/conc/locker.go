package conc

import "sync"

// RWLocker is a reader/writer mutual exclusion locker for concurrent access.
type RWLocker struct {
	// Indicates if this locker is enabled for synchronizing access.
	sync bool

	// Mutex used for synchronizing.
	mu sync.RWMutex
}

func (s *RWLocker) Read() func() {
	if !s.sync {
		return func() {}
	}

	s.mu.RLock()
	return s.mu.RUnlock
}

func (s *RWLocker) Write() func() {
	if !s.sync {
		return func() {}
	}

	s.mu.Lock()
	return s.mu.Unlock
}

// Synchronize sets the synchronization flag for this locker.
// If enable is true, it enables synchronization, meaning that
// the locker will synchronize concurrent access using locks.
// If enable is false, it disables synchronization, allowing
// concurrent access without additional synchronization.
func (s *RWLocker) Synchronize(enable bool) {
	s.sync = enable
}

// IsSynchronized returns true if the locker is configured for synchronization,
// meaning that concurrent access is synchronized using locks.
// Returns false if synchronization is disabled,
// allowing concurrent access without additional synchronization.
func (s *RWLocker) IsSynchronized() bool {
	return s.sync
}
