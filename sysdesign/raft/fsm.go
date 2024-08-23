package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"
)

type valueFSM struct {
	mtx   sync.RWMutex
	value string
}

var _ raft.FSM = (*valueFSM)(nil)

func (f *valueFSM) GetValue() string {
	f.mtx.RLock()
	defer f.mtx.RUnlock()
	return f.value
}

// Apply applies the log entry to the FSM's state.
func (f *valueFSM) Apply(l *raft.Log) interface{} {
	// Decode the command from the log entry
	var command SetValueCommand
	if err := json.Unmarshal(l.Data, &command); err != nil {
		return fmt.Errorf("failed to unmarshal command: %v", err)
	}

	// Acquire lock and update the state
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.value = command.Value

	return nil
}

// Snapshot creates a snapshot of the FSM's state.
func (f *valueFSM) Snapshot() (raft.FSMSnapshot, error) {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(f.value); err != nil {
		return nil, err
	}

	return &valueFSMSnapshot{buf: buf}, nil
}

// Restore restores the FSM's state from a snapshot.
func (f *valueFSM) Restore(snapshot io.ReadCloser) error {
	defer snapshot.Close()

	var restoredValue string
	if err := gob.NewDecoder(snapshot).Decode(&restoredValue); err != nil {
		return err
	}

	f.mtx.Lock()
	defer f.mtx.Unlock()

	f.value = restoredValue
	return nil
}

// valueFSMSnapshot is an implementation of raft.FSMSnapshot.
type valueFSMSnapshot struct {
	buf bytes.Buffer
}

// Persist implements raft.FSMSnapshot.
func (s *valueFSMSnapshot) Persist(sink raft.SnapshotSink) error {
	defer sink.Close()

	_, err := io.Copy(sink, &s.buf)
	if err != nil {
		return err
	}

	return nil
}

// Release implements raft.FSMSnapshot.
func (s *valueFSMSnapshot) Release() {

}
