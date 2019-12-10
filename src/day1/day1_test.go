package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	var initial int64 = 12
	var ans int64 = 2
	total := CalculateFuel(initial)

	if total != ans {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = 14
	ans = 2
	total = CalculateFuel(initial)

	if total != ans {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = 1969
	ans = 654
	total = CalculateFuel(initial)

	if total != ans {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = 100756
	ans = 33583
	total = CalculateFuel(initial)

	if total != ans {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = 14
	ans = 2
	total = CalculateFuelRepeat(initial)

	if total != ans {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}

	initial = 1969
	ans = 966
	total = CalculateFuelRepeat(initial)

	if total != 966 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, ans)
	}
}
