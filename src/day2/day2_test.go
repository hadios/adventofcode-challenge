package main

import (
	"testing"
)

func Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCalculateFuel(t *testing.T) {
	var initial = []int64{1, 0, 0, 0, 99}
	var ans = []int64{2, 0, 0, 0, 99}
	total := executeOpsCommand(initial)

	if !Equal(total, ans) {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []int64{2, 3, 0, 3, 99}
	ans = []int64{2, 3, 0, 6, 99}
	total = executeOpsCommand(initial)

	if !Equal(total, ans) {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []int64{2, 4, 4, 5, 99, 0}
	ans = []int64{2, 4, 4, 5, 99, 9801}
	total = executeOpsCommand(initial)

	if !Equal(total, ans) {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = []int64{1, 1, 1, 4, 99, 5, 6, 0, 99}
	ans = []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}
	total = executeOpsCommand(initial)

	if !Equal(total, ans) {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}
}
