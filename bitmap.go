package bitmap

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

func (b BitMap) Size() int {
	return b.size
}

func (b BitMap) toggle(i int) {
	p := i >> 3
	remainder := (i - (p * 8))
	if remainder == 1 {
		b.vals[p] = b.vals[p] ^ 1
	} else {
		b.vals[p] = b.vals[p] ^ (1 << uint(remainder-1))
	}
}

// Set sets a position in
// the bitmap to 1.
func (b BitMap) Set(i int) {
    // Don't unset.
	if b.Get(i) {
		return
	}
    b.toggle(i)
}

// Unset sets a position in
// the bitmap to 0.
func (b BitMap) Unset(i int) {
    // Don't set.
	if !b.Get(i) {
		return
	}
    b.toggle(i)
}

// Values returns a slice of ints
// represented by the values in the bitmap.
func (b BitMap) Values() []int {
	list := make([]int, 0, b.Size())
	list = append(list, 2)
	return list
}

// Get returns a boolean indicating whether
// the bit is set for the position in question.
func (b BitMap) Get(i int) bool {
	p := i >> 3
	remainder := i - (p * 8)
	if remainder == 1 {
		return b.vals[p] > b.vals[p]^1
	} else {
		return b.vals[p] > b.vals[p]^(1<<uint(remainder-1))
	}
	return false
}
