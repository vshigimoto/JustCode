package main

import (
	"testing"
)

func TestCompareTwoSlices(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	if compareTwoSlices(a, b) == false {
		t.Errorf("Slices not equal")
	}
	// if unorder
	c := []int{1, 3, 2, 4}
	if compareTwoSlices(a, c) == false {
		t.Errorf("Slices not equal")
	}
	// if emptySlice
	emptySlice := []int{}
	if compareTwoSlices(a, emptySlice) == false {
		t.Errorf("Slices not equal")
	}
	// if different length
	d := []int{1, 2, 3}
	if compareTwoSlices(a, d) == false {
		t.Errorf("Slices not equal")
	}
	// if different slices
	e := []int{5, 6, 7, 8}
	if compareTwoSlices(a, e) == false {
		t.Errorf("Slices not equal")
	}
	// one more test
	f := []int{1, 1, 1, 1}
	if compareTwoSlices(a, f) == false {
		t.Errorf("Slices not equal")
	}
}
