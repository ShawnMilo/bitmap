/*
Test the bitmap package.
*/

package bitmap_test

import (
	"bitmap"
	"testing"
    "fmt"
)

// TestCreate proves we can create a
// bitmap and its size is as set.
func TestCreate(t *testing.T) {
	for size := range []int{1, 13, 27, 66} {
		b := bitmap.New(size)
		if b.Size() != size {
			t.Error("size doesn't match")
		}
	}
}

func TestSet(t *testing.T) {
	b := bitmap.New(10)
	b.Set(2)
	if !slicesEqual(b.Values(), []int{2}) {
		t.Error("values do not match")
	}
}

func TestGet(t *testing.T) {
	b := bitmap.New(50)
	b.Set(2)
	if !b.Get(2) {
		t.Error("expected true, got false")
	}
	if b.Get(3) {
		t.Error("expected false, got true")
	}
	b.Set(3)
	if !b.Get(3) {
		t.Error("expected true, got false")
	}
    // Larger number (to prove indexing past the first
    // byte works.
	if b.Get(42) {
		t.Error("expected false, got true")
	}
	b.Set(42)
	if !b.Get(42) {
		t.Error("expected true, got false")
	}
}

// Setting a value twice should not unset it.
func TestSetTwice(t *testing.T) {
	b := bitmap.New(10)
	b.Set(2)
	if !b.Get(2){
		t.Error("expected it to be set")
	}
	b.Set(2)
	if !b.Get(2){
		t.Error("expected it to still be set")
	}
}


// Unset a value.
func TestSetUnset(t *testing.T) {
	b := bitmap.New(10)
	b.Set(2)
	if !b.Get(2){
		t.Error("expected it to be set")
	}
	b.Unset(2)
	if b.Get(2){
		t.Error("expected it to be unset")
	}
}

func ExampleBitmap () {
    b := bitmap.New(10)
    b.Set(2)
    fmt.Printf("2 in bitmap: %v. 7 in bitmap: %v.\n", b.Get(2), b.Get(7))
    // Output: 2 in bitmap: true. 7 in bitmap: false.
}
