/*
Utilities for bitmap tests and tests for those utilities.
*/
package bitmap_test

import (
	"testing"
)

// The slices here are equal.
var equalSlices = [][2][]int{
	[2][]int{[]int{2}, []int{2}},             // one value
	[2][]int{[]int{13, 17}, []int{13, 17}},   // two values
	[2][]int{[]int{1, 2, 3}, []int{1, 2, 3}}, // three values
	[2][]int{[]int{3, 1, 2}, []int{1, 2, 3}}, // three, mixed order
	[2][]int{[]int{13, 17}, []int{17, 13}},   // two, mixed order
}

// The slices here are not equal.
var inequalSlices = [][2][]int{
	[2][]int{[]int{2}, []int{2, 3}},                // wrong number of values on right
	[2][]int{[]int{1, 2, 3}, []int{1, 2}},          // wrong number on left
	[2][]int{[]int{9}, []int{7}},                   // simple
	[2][]int{[]int{13, 17}, []int{13, 18}},         // simple, multiple values
	[2][]int{[]int{21, 13, 17}, []int{13, 17}},     // Same values, extra on left
	[2][]int{[]int{21, 13}, []int{13, 17, 21}},     // Same values, extra on right
	[2][]int{[]int{13, 21}, []int{13, 17, 21}},     // Same values, but duplicate
}

// slicesEqual accepts two slices and returns a boolean
// indicating whether they are equal.
// Intentionally not implementing a sort, so this is
// a bit brute-force, but the amount of test data is small.
//
// Note: This function DOES return a false positive for a sample
// such as []int{1, 1, 2} == []int{1, 2, 2} because the bitmap can
// not have duplicate values.
func slicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, val := range s1 {
		found := false
		for _, x := range s2 {
			if x == val {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// TestEqual tests slicesEqual.
func TestEqual(t *testing.T) {
	for _, vals := range equalSlices {
		if !slicesEqual(vals[0], vals[1]) {
			t.Error("equal slices not found to be equal")
		}
	}
}

// TestNotEqual can't prove a negative, but can ensure
// that it doesn't always return true.
func TestNotEqual(t *testing.T) {
	for _, vals := range inequalSlices {
		if slicesEqual(vals[0], vals[1]) {
			t.Error("false positive found")
		}
	}
}
