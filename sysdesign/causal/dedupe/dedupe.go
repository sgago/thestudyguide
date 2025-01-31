// Package dedupe manages duplicate messages.
package dedupe

import (
	"sgago/thestudyguide-causal/lock"
	"time"
)

// Dedupe is a struct for managing duplicate messages.
type Dedupe[T comparable] struct {
	*lock.Lock

	seen map[T]time.Time
	ttl  time.Duration
}

// New creates a new instance of Deduplicator with the given time-to-live (TTL) duration.
func New[T comparable](ttl time.Duration) *Dedupe[T] {
	return &Dedupe[T]{
		Lock: lock.NewRW(),
		seen: make(map[T]time.Time),
		ttl:  ttl,
	}
}

// Seen checks if the given key has been seen before.
func (d *Dedupe[T]) Seen(key T) bool {
	defer d.Read()()

	_, ok := d.seen[key]
	return ok
}

// Mark marks the given key as seen and returns true if it was not seen before.
func (d *Dedupe[T]) Mark(key T) bool {
	defer d.Write()()

	var ok bool

	if _, ok = d.seen[key]; !ok {
		d.seen[key] = time.Now()
	}

	return !ok
}

// Delete removes the given key from the deduplicator.
func (d *Dedupe[T]) Delete(key T) {
	defer d.Write()()
	delete(d.seen, key)
}

// Clear clears all the keys from the deduplicator.
func (d *Dedupe[T]) Clear() {
	defer d.Write()()
	d.seen = make(map[T]time.Time)
}

// removeExpired periodically removes expired messages from the deduplicator's map.
func (d *Dedupe[T]) removeExpired() {
	for {
		time.Sleep(d.ttl / 2) // Cleanup runs half as often as the TTL duration

		d.Lock.Write()
		for id, timestamp := range d.seen {
			if time.Since(timestamp) > d.ttl {
				delete(d.seen, id)
			}
		}
		d.Lock.Unlock()
	}
}
