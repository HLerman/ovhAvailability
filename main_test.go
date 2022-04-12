package main

import "testing"

func TestIsIndex(t *testing.T) {
	v := IsIndex([]string{"element01", "element02", "element03"}, "element02")
	if v != 1 {
		t.Errorf("index is invalid, got: %d, want: %d.", v, 1)
	}

	v = IsIndex([]string{"element01", "element02", "element03"}, "element03")
	if v != 2 {
		t.Errorf("index is invalid, got: %d, want: %d.", v, 2)
	}

	v = IsIndex([]string{"element01", "element02", "element03"}, "element04")
	if v != 3 {
		t.Errorf("index is invalid, got: %d, want: %d.", v, 3)
	}

	v = IsIndex([]int{2, 6, 1}, 6)
	if v != 1 {
		t.Errorf("index is invalid, got: %d, want: %d.", v, 1)
	}
}
