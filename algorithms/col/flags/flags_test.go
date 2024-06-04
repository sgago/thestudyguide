package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitFlags(t *testing.T) {
	flags := New32(64)
	assert.Zero(t, flags.Ones())

	flags.Set(0, true)
	assert.True(t, flags.Get(0))
	assert.False(t, flags.Get(1))

	flags.Set(65, true)
	assert.True(t, flags.Get(65))
	assert.False(t, flags.Get(63))
	assert.False(t, flags.Get(64))

	assert.Equal(t, 2, flags.Ones())
	assert.Equal(t, 94, flags.Zeros())
}
