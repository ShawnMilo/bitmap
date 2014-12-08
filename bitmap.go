package bitmap

import "errors"

var ErrOutOfRange = errors.New("Value out of range.")

// BitMap is a struct containing a slice of bytes,
// being used as a bitmap.
type BitMap struct {
	size int
	vals []byte
}

// New returns a BitMap. It requires a size. A bitmap with a size of
// eight or less will be one byte in size, and so on.
func New(s int) BitMap {
	l := s/8 + 1
	return BitMap{size: s, vals: make([]byte, l, l)}
}

// Size returns the size of a bitmap. This is the number
// of bits.
func (b BitMap) Size() int {
	return b.size
}

// checkRange returns an error if the position
// passed is not allowed.
func (b BitMap) checkRange(i int) error {
	if i > b.Size() {
		return ErrOutOfRange
	}
	if i < 1 {
		return ErrOutOfRange
	}
	return nil
}

// For internal use; drives Set and Unset.
func (b BitMap) toggle(i int) {
	// Position of the byte in b.vals.
	p := i >> 3
	// Position of the bit in the byte.
	remainder := i - (p * 8)
	// Toggle the bit.
	if remainder == 1 {
		b.vals[p] = b.vals[p] ^ 1
	} else {
		b.vals[p] = b.vals[p] ^ (1 << uint(remainder-1))
	}
}

// Set sets a position in
// the bitmap to 1.
func (b BitMap) Set(i int) error {
	if x := b.checkRange(i); x != nil {
		return x
	}
	// Don't unset.
	val, err := b.Get(i)
	if err != nil {
		return err
	}
	if val {
		return nil
	}
	b.toggle(i)
	return nil
}

// Unset sets a position in
// the bitmap to 0.
func (b BitMap) Unset(i int) error {
	// Don't set.
	val, err := b.Get(i)
	if err != nil {
		return err
	}
	if val {
		b.toggle(i)
	}
	return nil
}

// Values returns a slice of ints
// represented by the values in the bitmap.
func (b BitMap) Values() ([]int, error) {
	list := make([]int, 0, b.Size())
	for i := 1; i <= b.Size(); i++ {
		val, err := b.Get(i)
		if err != nil {
			return nil, err
		}
		if val {
			list = append(list, i)
		}
	}
	return list, nil
}

// Get returns a boolean indicating whether
// the bit is set for the position in question.
func (b BitMap) Get(i int) (bool, error) {
	if x := b.checkRange(i); x != nil {
		return false, x
	}
	p := i >> 3
	remainder := i - (p * 8)
	if remainder == 1 {
		return b.vals[p] > b.vals[p]^1, nil
	}
	return b.vals[p] > b.vals[p]^(1<<uint(remainder-1)), nil
}
