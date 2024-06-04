// Package flags implements boolean flagging via a
// slice of unsigned integers and bitmasking.
// It is intended to replace arrays or slices of booleans.
// Useful for graph searches to avoid cycles or combinatorial
// problems with use-once-values.
package flags

import "math"

type UInt interface {
	uint8 | uint16 | uint32 | uint64
}

type Flags interface {
	Get(pos int) bool

	Set(pos int, val bool)
	True(pos int)
	False(pos int)

	Toggle(pos int)
	ToggleAll()

	Ones() int
	Zeros() int
}

type BitFlags[T UInt] struct {
	bits   []T
	size   int
	maxVal T
}

var _ Flags = (*BitFlags[uint8])(nil)

func New(bits int) Flags {
	if bits <= 8 {
		return newFlags[uint8](bits, 8, math.MaxUint8)
	} else if bits <= 16 {
		return newFlags[uint16](bits, 16, math.MaxUint16)
	} else if bits <= 32 {
		return newFlags[uint32](bits, 32, math.MaxUint32)
	}

	return newFlags[uint64](bits, 64, math.MaxUint64)
}

func New8(bits int) BitFlags[uint8] {
	return *newFlags[uint8](bits, 8, math.MaxUint8)
}

func New16(bits int) BitFlags[uint16] {
	return *newFlags[uint16](bits, 16, math.MaxUint16)
}

func New32(bits int) BitFlags[uint32] {
	return *newFlags[uint32](bits, 32, math.MaxUint32)
}

func New64(bits int) BitFlags[uint64] {
	return *newFlags[uint64](bits, 64, math.MaxUint64)
}

func newFlags[T UInt](bits, size int, maxVal T) *BitFlags[T] {
	return &BitFlags[T]{
		bits:   make([]T, max(1, bits)/size),
		size:   size,
		maxVal: maxVal,
	}
}

func (f *BitFlags[T]) Get(pos int) bool {
	idx := pos / f.size
	bit := pos % f.size

	if idx >= len(f.bits) {
		return false
	}

	return get(f.bits[idx], bit)
}

func (f *BitFlags[T]) Set(pos int, val bool) {
	idx := pos / f.size
	bit := pos % f.size

	if idx >= len(f.bits) {
		f.bits = append(f.bits, make([]T, idx-len(f.bits)+1)...)
	}

	if val {
		f.bits[idx] = set(f.bits[idx], bit)
	} else {
		f.bits[idx] = clear(f.bits[idx], bit)
	}
}

func (f *BitFlags[T]) True(pos int) {
	f.Set(pos, true)
}

func (f *BitFlags[T]) False(pos int) {
	f.Set(pos, false)
}

func (f *BitFlags[T]) Toggle(pos int) {
	idx := pos / f.size
	bit := pos % f.size

	f.bits[idx] = toggle(f.bits[idx], bit)
}

func (f *BitFlags[T]) ToggleAll() {
	for i, n := range f.bits {
		f.bits[i] = toggleAll(n, f.maxVal)
	}
}

func (f *BitFlags[T]) Ones() int {
	sum := 0

	for _, n := range f.bits {
		sum += ones(n)
	}

	return sum
}

func (f *BitFlags[T]) Zeros() int {
	return (len(f.bits) * f.size) - f.Ones()
}

func get[T UInt](n T, pos int) bool {
	return (n & (1 << pos)) > 0
}

func set[T UInt](n T, pos int) T {
	return n | (1 << pos)
}

func clear[T UInt](n T, pos int) T {
	return n & ^(1 << pos)
}

func ones[T UInt](n T) int {
	cnt := 0

	for ; n > 0; cnt++ {
		n = n & (n - 1)
	}

	return cnt
}

func toggle[T UInt](n T, pos int) T {
	return n ^ (1 << pos)
}

func toggleAll[T UInt](n T, max T) T {
	return n ^ max
}
