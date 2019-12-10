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
	var initial int64 = 111111
	var ans = true
	var total = isIncreasingDigits(initial) && hasDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 223450
	ans = false
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 123789
	ans = false
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 112233
	ans = true
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 123444
	ans = false
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 111122
	ans = true
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 123455
	ans = true
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 222333
	ans = false
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}

	initial = 188999
	ans = true
	total = isIncreasingDigits(initial) && hasDoubleDigits(initial) && hasNoLargerDoubleDigits(initial)

	if total != ans {
		t.Errorf("Result was incorrect, got: %t, want: %t.", total, ans)
	}
}
