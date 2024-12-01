package util

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func AssertEqualArrays(t *testing.T, a, b []int) {
	if !Equal(a, b) {
		t.Errorf("got %v want %v", a, b)
	}
}

func AssertEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("got %v want %v", a, b)
	}
}
