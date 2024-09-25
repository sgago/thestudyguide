// Package flags implements boolean flagging via a slice of unsigned integers and bitmasking.
// It is intended to replace arrays or slices of booleans.
// Useful for graph searches to avoid cycles or combinatorial problems with use-once-values.
package flags

import "math"

// UInt is an interface that represents an unsigned integer type.
// It can be implemented by types uint8, uint16, uint32, and uint64.
type UInt interface {
	uint8 | uint16 | uint32 | uint64
}

// Flags represents a collection of boolean values that can be manipulated individually or as a whole.
type Flags interface {
	// Get retrieves the boolean value at the specified position.
	Get(pos int) bool

	// Set sets the boolean value at the specified position.
	Set(pos int, val bool)

	// True sets the boolean value at the specified position to true.
	True(pos int)

	// False sets the boolean value at the specified position to false.
	False(pos int)

	// Toggle inverts the boolean value at the specified position.
	Toggle(pos int)

	// ToggleAll inverts all boolean values in the collection.
	ToggleAll()

	// Ones returns the number of true values in the collection.
	Ones() int

	// Zeros returns the number of false values in the collection.
	Zeros() int
}

// BitFlags represents a collection of bit flags.
type BitFlags[T UInt] struct {
	bits   []T // The underlying array of bits.
	size   int // The size of the bit flags collection.
	maxVal T   // The maximum value that can be represented by the bit flags.
}

var _ Flags = (*BitFlags[uint8])(nil)

// New creates a new Flags instance with the specified number of bits.
// It returns a Flags object that can be used to manipulate and query individual flags.
// The bits parameter specifies the number of bits to allocate for the flags.
// If bits is less than or equal to 8, a Flags object with 8 bits will be created.
// If bits is less than or equal to 16, a Flags object with 16 bits will be created.
// If bits is less than or equal to 32, a Flags object with 32 bits will be created.
// If bits is greater than 32, a Flags object with 64 bits will be created.
// The maximum value for each type of Flags object is set according to the maximum value of the corresponding unsigned integer type.
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

// New8 creates a new instance of BitFlags with a specified number of bits.
// It returns a BitFlags[uint8] object.
// The 'bits' parameter specifies the number of bits for the BitFlags object.
// The BitFlags object is initialized with 8 bits and a maximum value of math.MaxUint8.
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

// Get returns the value of the bit at the specified position in the BitFlags.
// It takes an integer position as input and returns a boolean value.
// If the position is out of range, it returns false.
func (f *BitFlags[T]) Get(pos int) bool {
	idx := pos / f.size
	bit := pos % f.size

	if idx >= len(f.bits) {
		return false
	}

	return get(f.bits[idx], bit)
}

// Set sets the value of the bit at the specified position in the BitFlags.
// It takes an integer position and sets a boolean value at that position.
// If the position is out of range, Set will extend the BitFlags to accommodate the new position.
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

// True sets the bit at the specified position to true.
func (f *BitFlags[T]) True(pos int) {
	f.Set(pos, true)
}

// False sets the bit at the specified position to false.
func (f *BitFlags[T]) False(pos int) {
	f.Set(pos, false)
}

// Toggle toggles (a.k.a. flips or inverts) the bit at the specified position in the BitFlags.
// It takes the position of the bit to toggle as input.
// The position is zero-based.
// The function updates the BitFlags by toggling the bit at the specified position.
func (f *BitFlags[T]) Toggle(pos int) {
	idx := pos / f.size
	bit := pos % f.size

	f.bits[idx] = toggle(f.bits[idx], bit)
}

// ToggleAll toggles (a.k.a. flips or inverts) all bits in the BitFlags.
// It flips all bits from 0 to the maximum value of the BitFlags.
// The function updates the BitFlags by toggling all bits.
func (f *BitFlags[T]) ToggleAll() {
	for i, n := range f.bits {
		f.bits[i] = toggleAll(n, f.maxVal)
	}
}

// Ones returns the number of true values in the BitFlags.
func (f *BitFlags[T]) Ones() int {
	sum := 0

	for _, n := range f.bits {
		sum += ones(n)
	}

	return sum
}

// Zeros returns the number of false values in the BitFlags.
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
