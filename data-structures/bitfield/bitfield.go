package bitfield

// Bitfield is a golang implementation of a bitfield
type Bitfield []byte

// New returns a new Bitfield with minimum of n bits
func New(n int) Bitfield {
	// calculate bytes required to store n bits
	n = 1 + ((n - 1) / 8)
	return Bitfield(make([]byte, n))
}

// Clear sets the bit that represents i to 0
func (b Bitfield) Clear(i int) {
	index := i / 8
	offset := i % 8
	b[index] &= ^(1 << offset)
}

// Contains returns true if the bit that represents i is set to 1
func (b Bitfield) Contains(i int) bool {
	index := i / 8
	offset := i % 8
	if index > len(b) {
		return false
	}
	return (b[index] & (1 << offset)) != 0
}

// Reset sets all bits to 0
func (b Bitfield) Reset() {
	for i := range b {
		b[i] = 0
	}
}

// Set sets the bit that represents i to 1
func (b Bitfield) Set(i int) {
	index := i / 8
	offset := i % 8
	b[index] |= (1 << offset)
}
