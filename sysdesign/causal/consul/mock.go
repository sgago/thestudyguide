package consul

import (
	"fmt"
	"sgago/thestudyguide-causal/envcfg"
)

type MockCli struct{}

var _ Client = (*MockCli)(nil)

func NewMock() *MockCli {
	return &MockCli{}
}

// Register implements Client.
func (m *MockCli) Register() error {
	return nil
}

// StartTTL implements Client.
func (m *MockCli) StartTTL() {
	// No op
}

// StopTTL implements Client.
func (m *MockCli) StopTTL() {
	// No op
}

// Self implements Client.
func (consul *MockCli) Self() string {
	return fmt.Sprintf("http://%s:%d", envcfg.HostName(), envcfg.Port())
}

// Services implements Client.
func (m *MockCli) Services() []string {
	return []string{m.Self()}
}
